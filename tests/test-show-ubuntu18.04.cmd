footloose config create --override --config %testName.footloose --name %testName --key %testName-key --image ubuntu18.04
%defer footloose delete --config %testName.footloose
footloose create --config %testName.footloose
footloose delete --config %testName.footloose
%out footloose show --config %testName.footloose
%out footloose show -o json --config %testName.footloose
