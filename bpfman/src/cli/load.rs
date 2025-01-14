// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of bpfman

use std::collections::HashMap;

use anyhow::bail;
use bpfman_api::{TcProceedOn, XdpProceedOn};

use crate::{
    bpf::BpfManager,
    cli::{
        args::{GlobalArg, LoadCommands, LoadFileArgs, LoadImageArgs, LoadSubcommand},
        table::ProgTable,
    },
    command::{
        KprobeProgram, Program, ProgramData, TcProgram, TracepointProgram, UprobeProgram,
        XdpProgram,
    },
};

impl LoadSubcommand {
    pub(crate) async fn execute(&self, bpf_manager: &mut BpfManager) -> anyhow::Result<()> {
        match self {
            LoadSubcommand::File(l) => execute_load_file(bpf_manager, l).await,
            LoadSubcommand::Image(l) => execute_load_image(bpf_manager, l).await,
        }
    }
}

pub(crate) async fn execute_load_file(
    bpf_manager: &mut BpfManager,
    args: &LoadFileArgs,
) -> anyhow::Result<()> {
    let bytecode_source = crate::command::Location::File(args.path.clone());

    let data = ProgramData::new_pre_load(
        bytecode_source,
        args.name.clone(),
        args.metadata
            .clone()
            .unwrap_or_default()
            .iter()
            .map(|(k, v)| (k.to_owned(), v.to_owned()))
            .collect(),
        parse_global(&args.global),
        args.map_owner_id,
    )?;

    let program = bpf_manager
        .add_program(args.command.get_program(data)?)
        .await?;

    ProgTable::new_get_bpfman(&Some((&program).try_into()?))?.print();
    ProgTable::new_get_unsupported(&Some((&program).try_into()?))?.print();
    Ok(())
}

pub(crate) async fn execute_load_image(
    bpf_manager: &mut BpfManager,
    args: &LoadImageArgs,
) -> anyhow::Result<()> {
    let bytecode_source = crate::command::Location::Image((&args.pull_args).try_into()?);

    let data = ProgramData::new_pre_load(
        bytecode_source,
        args.name.clone(),
        args.metadata
            .clone()
            .unwrap_or_default()
            .iter()
            .map(|(k, v)| (k.to_owned(), v.to_owned()))
            .collect(),
        parse_global(&args.global),
        args.map_owner_id,
    )?;

    let program = bpf_manager
        .add_program(args.command.get_program(data)?)
        .await?;

    ProgTable::new_get_bpfman(&Some((&program).try_into()?))?.print();
    ProgTable::new_get_unsupported(&Some((&program).try_into()?))?.print();
    Ok(())
}

impl LoadCommands {
    pub(crate) fn get_program(&self, data: ProgramData) -> Result<Program, anyhow::Error> {
        match self {
            LoadCommands::Xdp {
                iface,
                priority,
                proceed_on,
            } => {
                let proc_on = match XdpProceedOn::from_strings(proceed_on) {
                    Ok(p) => p,
                    Err(e) => bail!("error parsing proceed_on {e}"),
                };
                Ok(Program::Xdp(XdpProgram::new(
                    data,
                    *priority,
                    iface.to_string(),
                    XdpProceedOn::from_int32s(proc_on.as_action_vec())?,
                )?))
            }
            LoadCommands::Tc {
                direction,
                iface,
                priority,
                proceed_on,
            } => {
                match direction.as_str() {
                    "ingress" | "egress" => (),
                    other => bail!("{} is not a valid direction", other),
                };
                let proc_on = match TcProceedOn::from_strings(proceed_on) {
                    Ok(p) => p,
                    Err(e) => bail!("error parsing proceed_on {e}"),
                };
                Ok(Program::Tc(TcProgram::new(
                    data,
                    *priority,
                    iface.to_string(),
                    proc_on,
                    direction.to_string().try_into()?,
                )?))
            }
            LoadCommands::Tracepoint { tracepoint } => Ok(Program::Tracepoint(
                TracepointProgram::new(data, tracepoint.to_string())?,
            )),
            LoadCommands::Kprobe {
                fn_name,
                offset,
                retprobe,
                container_pid,
            } => {
                if container_pid.is_some() {
                    bail!("kprobe container option not supported yet");
                }
                let offset = offset.unwrap_or(0);
                Ok(Program::Kprobe(KprobeProgram::new(
                    data,
                    fn_name.to_string(),
                    offset,
                    *retprobe,
                    None,
                )?))
            }
            LoadCommands::Uprobe {
                fn_name,
                offset,
                target,
                retprobe,
                pid,
                container_pid,
            } => {
                let offset = offset.unwrap_or(0);
                Ok(Program::Uprobe(UprobeProgram::new(
                    data,
                    fn_name.clone(),
                    offset,
                    target.to_string(),
                    *retprobe,
                    *pid,
                    *container_pid,
                )?))
            }
        }
    }
}

fn parse_global(global: &Option<Vec<GlobalArg>>) -> HashMap<String, Vec<u8>> {
    let mut global_data: HashMap<String, Vec<u8>> = HashMap::new();

    if let Some(global) = global {
        for g in global.iter() {
            global_data.insert(g.name.to_string(), g.value.clone());
        }
    }
    global_data
}
