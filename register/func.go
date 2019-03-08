package register

import "reflect"

func (reg *SRegister) GetApp() string {
	return config.App
}

func (reg *SRegister) GetName() string {
	return config.Name
}
func (reg *SRegister) GetVersion() string {
	return IntToStr(config.Version)
}
func (reg *SRegister) GetKey() string {
	return config.Key
}

func (reg *SRegister) SetRegSuccess(isReg bool) {
	reg.RegSuccess = isReg
}

func (reg *SRegister) SetServiceData(data map[string]map[string]string) {
	reg.ServiceData = data
}

func (reg *SRegister) GetFunctions() []string {
	return GetHproseAddedFunc()
}

func (reg *SRegister) CloseClient() {
	reg.Conn.Close()
}

func (reg *SRegister) Info(arg0 interface{}, args ...interface{}) {
	reg.Logger.Info(arg0, args...)
}

func (reg *SRegister) Debug(arg0 interface{}, args ...interface{}) {
	reg.Logger.Debug(arg0, args...)
}

func (reg *SRegister) Warn(arg0 interface{}, args ...interface{}) {
	reg.Logger.Warn(arg0, args...)
}

func (reg *SRegister) Error(arg0 interface{}, args ...interface{}) {
	reg.Logger.Error(arg0, args...)
}

// 调取rpc服务
func (reg *SRegister) RpcCall(name string, args []reflect.Value, namespace string, cfg map[string]uint32) interface{} {
	var serviceId uint32
	if _, ok := cfg["serviceId"]; ok {
		serviceId = cfg["serviceId"]
	}
	var adminId uint32
	if _, ok := cfg["adminId"]; ok {
		adminId = cfg["adminId"]
	}

	rpc := NewRpcClient(RpcClient{
		NameSpace: namespace,
		ServiceId: serviceId,
		AdminId: adminId,
	})

	return rpc.Call(name, args)
}
