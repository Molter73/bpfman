kubectl get bpfprogramconfigs -o custom-columns="NAME":.metadata.name,"TYPE":.spec.type,"SECNAME":.spec.name,"STATUS":.status.conditions[-1:].type,"INTERFACE":.spec.attachpoint.networkmultiattach.interface,"PRIORITY":.spec.attachpoint.networkmultiattach.priority,"DIRECTION":.spec.attachpoint.networkmultiattach.direction,"TRACEPOINT":.spec.attachpoint.singleattach.name

