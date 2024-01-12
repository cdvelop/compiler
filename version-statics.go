package compiler

func (c Compiler) versionStatics() string {
	return "?=" + c.Config.AppVersion()
}
