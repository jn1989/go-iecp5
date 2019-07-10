package asdu

import "strconv"

// about data unit identification 应用服务数据单元 - 数据单元标识符

// TypeID is the ASDU type identification.
type TypeID uint8

// The standard ASDU type identification.
// M for monitored information
// C for control information
// P for parameter
// F for file transfer.
// <0> 未用
// <1..127> 标准定义 - 兼容
// <128..135> 为路由报文保留 - 专用
// <136..255> 特殊应用 - 专用
// NOTE: 信息对象带或不带时标由标识符类型的不同序列来区别
const (
	_ TypeID = iota // 0: not defined
	// 在监视方向上的过程信息 <0..44>
	M_SP_NA_1 // single-point information, 单点信息
	M_SP_TA_1 // single-point information with time tag, 单点信息-带时标
	M_DP_NA_1 // double-point information, 双点信息
	M_DP_TA_1 // double-point information with time tag, 双点信息-带时标
	M_ST_NA_1 // step position information, 步位置信息
	M_ST_TA_1 // step position information with time tag, 步位置信息-带时标
	M_BO_NA_1 // bitstring of 32 bit, 32位比特串
	M_BO_TA_1 // bitstring of 32 bit with time tag, 32位比特串-带时标
	M_ME_NA_1 // measured value, normalized value, 测量值，规一化值
	M_ME_TA_1 // measured value, normalized value with time tag, 测量值，规一化值-带时标
	M_ME_NB_1 // measured value, scaled value, 测量值，标度化值
	M_ME_TB_1 // measured value, scaled value with time tag, 测量值带时标，标度化值-带时标
	M_ME_NC_1 // measured value, short floating point number, 测量值，短浮点数
	M_ME_TC_1 // measured value, short floating point number with time tag, 测量值，短浮数-带时标
	M_IT_NA_1 // integrated totals, 累积量
	M_IT_TA_1 // integrated totals with time tag, 累积量带时标
	M_EP_TA_1 // event of protection equipment with time tag, 继电器保护设备事件-带时标
	M_EP_TB_1 // packed start events of protection equipment with time tag, 继电保护设备成组启动事件-带时标
	M_EP_TC_1 // packed output circuit information of protection equipment with time tag, 继电保护设备成组输出电路信息-带时标
	M_PS_NA_1 // packed single-point information with status change detection, 带变位检出的成组单点信息
	M_ME_ND_1 // measured value, normalized value without quality descriptor, 测量值,不带品质描述词的规一化值
	_         // 22: reserved for further compatible definitions
	_         // 23: reserved for further compatible definitions
	_         // 24: reserved for further compatible definitions
	_         // 25: reserved for further compatible definitions
	_         // 26: reserved for further compatible definitions
	_         // 27: reserved for further compatible definitions
	_         // 28: reserved for further compatible definitions
	_         // 29: reserved for further compatible definitions
	M_SP_TB_1 // single-point information with time tag CP56Time2a, 单点信息-带CP56Time2a
	M_DP_TB_1 // double-point information with time tag CP56Time2a, 双点信息-带CP56Time2a
	M_ST_TB_1 // step position information with time tag CP56Time2a, 步位置信息-带CP56Time2a
	M_BO_TB_1 // bitstring of 32 bits with time tag CP56Time2a, 32比特串-带CP56Time2a
	M_ME_TD_1 // measured value, normalized value with time tag CP56Time2a, 测量值,规一化值-带CP56Time2a
	M_ME_TE_1 // measured value, scaled value with time tag CP56Time2a, 测量值,标度化值-带CP56Time2a
	M_ME_TF_1 // measured value, short floating point number with time tag CP56Time2a, 测量值,短浮点数-带CP56Time2a
	M_IT_TB_1 // integrated totals with time tag CP56Time2a, 累积值-带CP56Time2a
	M_EP_TD_1 // event of protection equipment with time tag CP56Time2a, 继电保护装置事件-带CP56Time2a
	M_EP_TE_1 // packed start events of protection equipment with time tag CP56Time2a, 继电保护装置成组启动事件-带CP56Time2a
	M_EP_TF_1 // packed output circuit information of protection equipment with time tag CP56Time2a, 继电保护装置成组输出电路信息-带CP56Time2a
	S_IT_TC_1 // integrated totals containing time-tagged security statistics
	_         // 42: reserved for further compatible definitions
	_         // 43: reserved for further compatible definitions
	_         // 44: reserved for further compatible definitions
	// 在控制方向的过程信息 <45..69>
	C_SC_NA_1 // single command 单点命令
	C_DC_NA_1 // double command 双点命令
	C_RC_NA_1 // regulating step command 调节步命令
	C_SE_NA_1 // set-point command, normalized value 设定值命令，归一化值
	C_SE_NB_1 // set-point command, scaled value 设定值命令，规度化值
	C_SE_NC_1 // set-point command, short floating point number 设定值命令，短浮点数值
	C_BO_NA_1 // bitstring of 32 bits 23位比特串
	_         // 52: reserved for further compatible definitions
	_         // 53: reserved for further compatible definitions
	_         // 54: reserved for further compatible definitions
	_         // 55: reserved for further compatible definitions
	_         // 56: reserved for further compatible definitions
	_         // 57: reserved for further compatible definitions
	C_SC_TA_1 // single command with time tag CP56Time2a
	C_DC_TA_1 // double command with time tag CP56Time2a
	C_RC_TA_1 // regulating step command with time tag CP56Time2a
	C_SE_TA_1 // set-point command with time tag CP56Time2a, normalized value
	C_SE_TB_1 // set-point command with time tag CP56Time2a, scaled value
	C_SE_TC_1 // set-point command with time tag CP56Time2a, short floating point number
	C_BO_TA_1 // bitstring of 32-bit with time tag CP56Time2a
	_         // 65: reserved for further compatible definitions
	_         // 66: reserved for further compatible definitions
	_         // 67: reserved for further compatible definitions
	_         // 68: reserved for further compatible definitions
	_         // 69: reserved for further compatible definitions
	// 在监视方向的系统命令 <70..99>
	M_EI_NA_1 // end of initialization  初始化结束
	_         // 71: reserved for further compatible definitions
	_         // 72: reserved for further compatible definitions
	_         // 73: reserved for further compatible definitions
	_         // 74: reserved for further compatible definitions
	_         // 75: reserved for further compatible definitions
	_         // 76: reserved for further compatible definitions
	_         // 77: reserved for further compatible definitions
	_         // 78: reserved for further compatible definitions
	_         // 79: reserved for further compatible definitions
	_         // 80: reserved for further compatible definitions
	S_CH_NA_1 // authentication challenge
	S_RP_NA_1 // authentication reply
	S_AR_NA_1 // aggressive mode authentication request
	S_KR_NA_1 // session key status request
	S_KS_NA_1 // session key status
	S_KC_NA_1 // session key change
	S_ER_NA_1 // authentication error
	_         // 88: reserved for further compatible definitions
	_         // 89: reserved for further compatible definitions
	S_US_NA_1 // user status change
	S_UQ_NA_1 // update key change request
	S_UR_NA_1 // update key change reply
	S_UK_NA_1 // update key change — symetric
	S_UA_NA_1 // update key change — asymetric
	S_UC_NA_1 // update key change confirmation
	_         // 96: reserved for further compatible definitions
	_         // 97: reserved for further compatible definitions
	_         // 98: reserved for further compatible definitions
	_         // 99: reserved for further compatible definitions
	// 在控制方向的系统命令 <100..109>
	C_IC_NA_1 // interrogation command 总召唤
	C_CI_NA_1 // counter interrogation command 计数量召唤
	C_RD_NA_1 // read command 读命令
	C_CS_NA_1 // clock synchronization command 时钟同步命令
	C_TS_NA_1 // test command 测试命令
	C_RP_NA_1 // reset process command 复位进程命令
	C_CD_NA_1 // delay acquisition command 延时获得命令
	C_TS_TA_1 // test command with time tag CP56Time2a  带CP56Time2a的测试命令
	_         // 108: reserved for further compatible definitions
	_         // 109: reserved for further compatible definitions
	// 在控制方向的参数命令 <110..119>
	P_ME_NA_1 // parameter of measured value, normalized value 测量值参数,规一化值
	P_ME_NB_1 // parameter of measured value, scaled value 测量值参数,标度化值
	P_ME_NC_1 // parameter of measured value, short floating point number 测量值参数,短浮点数
	P_AC_NA_1 // parameter activation 参数激活
	_         // 114: reserved for further compatible definitions
	_         // 115: reserved for further compatible definitions
	_         // 116: reserved for further compatible definitions
	_         // 117: reserved for further compatible definitions
	_         // 118: reserved for further compatible definitions
	_         // 119: reserved for further compatible definitions
	// 文件传输 <120..127>
	F_FR_NA_1 // file ready  文件准备就绪
	F_SR_NA_1 // section ready 节准备就绪
	F_SC_NA_1 // call directory, select file, call file, call section 如唤目录，选择文件，召唤文件，召唤节
	F_LS_NA_1 // last section, last segment 最后的节，最后的段
	F_AF_NA_1 // ack file, ack section 认可文件，认可节
	F_SG_NA_1 // segment 段
	F_DR_TA_1 // directory 目录
	F_SC_NB_1 // QueryLog - request archive file (section 104) 查询日志
)

// infoObjSize maps the type identification (TypeID) to the serial octet size.
// Type extensions must register here.
var infoObjSize = [256]int{
	M_SP_NA_1: 1,
	M_SP_TA_1: 4,
	M_DP_NA_1: 1,
	M_DP_TA_1: 4,
	M_ST_NA_1: 2,
	M_ST_TA_1: 5,
	M_BO_NA_1: 5,
	M_BO_TA_1: 8,
	M_ME_NA_1: 3,
	M_ME_TA_1: 6,
	M_ME_NB_1: 3,
	M_ME_TB_1: 6,
	M_ME_NC_1: 5,
	M_ME_TC_1: 8,
	M_IT_NA_1: 5,
	M_IT_TA_1: 8,
	M_EP_TA_1: 6,
	M_EP_TB_1: 7,
	M_EP_TC_1: 7,
	M_PS_NA_1: 5,
	M_ME_ND_1: 2,

	M_SP_TB_1: 8,
	M_DP_TB_1: 8,
	M_ST_TB_1: 9,
	M_BO_TB_1: 12,
	M_ME_TD_1: 10,
	M_ME_TE_1: 10,
	M_ME_TF_1: 12,
	M_IT_TB_1: 12,
	M_EP_TD_1: 11,
	M_EP_TE_1: 11,
	M_EP_TF_1: 11,

	C_SC_NA_1: 1,
	C_DC_NA_1: 1,
	C_RC_NA_1: 1,
	C_SE_NA_1: 3,
	C_SE_NB_1: 3,
	C_SE_NC_1: 5,
	C_BO_NA_1: 4,

	M_EI_NA_1: 1,

	C_IC_NA_1: 1,
	C_CI_NA_1: 1,
	C_RD_NA_1: 0,
	C_CS_NA_1: 7,
	C_TS_NA_1: 2,
	C_RP_NA_1: 1,
	C_CD_NA_1: 2,

	P_ME_NA_1: 3,
	P_ME_NB_1: 3,
	P_ME_NC_1: 5,
	P_AC_NA_1: 1,

	F_FR_NA_1: 6,
	F_SR_NA_1: 7,
	F_SC_NA_1: 4,
	F_LS_NA_1: 5,
	F_AF_NA_1: 4,
	// F_SG_NA_1: 4 + variable,
	F_DR_TA_1: 13,
}

// GetInfoObjSize get the serial octet size of the type identification (TypeID).
func GetInfoObjSize(id TypeID) (int, error) {
	size := infoObjSize[id]
	if size == 0 {
		return 0, errTypeIdentifier
	}
	return size, nil
}

const (
	_TypeIDName0 = "M_SP_NA_1M_SP_TA_1M_DP_NA_1M_DP_TA_1M_ST_NA_1M_ST_TA_1M_BO_NA_1M_BO_TA_1M_ME_NA_1M_ME_TA_1M_ME_NB_1M_ME_TB_1M_ME_NC_1M_ME_TC_1M_IT_NA_1M_IT_TA_1M_EP_TA_1M_EP_TB_1M_EP_TC_1M_PS_NA_1M_ME_ND_1"
	_TypeIDName1 = "M_SP_TB_1M_DP_TB_1M_ST_TB_1M_BO_TB_1M_ME_TD_1M_ME_TE_1M_ME_TF_1M_IT_TB_1M_EP_TD_1M_EP_TE_1M_EP_TF_1S_IT_TC_1"
	_TypeIDName2 = "C_SC_NA_1C_DC_NA_1C_RC_NA_1C_SE_NA_1C_SE_NB_1C_SE_NC_1C_BO_NA_1"
	_TypeIDName3 = "C_SC_TA_1C_DC_TA_1C_RC_TA_1C_SE_TA_1C_SE_TB_1C_SE_TC_1C_BO_TA_1"
	_TypeIDName4 = "M_EI_NA_1"
	_TypeIDName5 = "S_CH_NA_1S_RP_NA_1S_AR_NA_1S_KR_NA_1S_KS_NA_1S_KC_NA_1S_ER_NA_1"
	_TypeIDName6 = "S_US_NA_1S_UQ_NA_1S_UR_NA_1S_UK_NA_1S_UA_NA_1S_UC_NA_1"
	_TypeIDName7 = "C_IC_NA_1C_CI_NA_1C_RD_NA_1C_CS_NA_1C_TS_NA_1C_RP_NA_1C_CD_NA_1C_TS_TA_1"
	_TypeIDName8 = "P_ME_NA_1P_ME_NB_1P_ME_NC_1P_AC_NA_1"
	_TypeIDName9 = "F_FR_NA_1F_SR_NA_1F_SC_NA_1F_LS_NA_1F_AF_NA_1F_SG_NA_1F_DR_TA_1F_SC_NB_1"
)

func (this TypeID) String() string {
	switch {
	case 1 <= this && this <= 21:
		this -= 1
		return _TypeIDName0[this*9 : 9*(this+1)]
	case 30 <= this && this <= 41:
		this -= 30
		return _TypeIDName1[this*9 : 9*(this+1)]
	case 45 <= this && this <= 51:
		this -= 45
		return _TypeIDName2[this*9 : 9*(this+1)]
	case 58 <= this && this <= 64:
		this -= 58
		return _TypeIDName3[this*9 : 9*(this+1)]
	case this == 70:
		return _TypeIDName4
	case 81 <= this && this <= 87:
		this -= 81
		return _TypeIDName5[this*9 : 9*(this+1)]
	case 90 <= this && this <= 95:
		this -= 90
		return _TypeIDName6[this*9 : 9*(this+1)]
	case 100 <= this && this <= 107:
		this -= 100
		return _TypeIDName7[this*9 : 9*(this+1)]
	case 110 <= this && this <= 113:
		this -= 110
		return _TypeIDName8[this*9 : 9*(this+1)]
	case 120 <= this && this <= 127:
		this -= 120
		return _TypeIDName9[this*9 : 9*(this+1)]
	default:
		return "TypeID(" + strconv.FormatInt(int64(this), 10) + ")"
	}
}

// Variable is variable structure qualifier
// number <0..127>, bit0 - bit6
// seq, bit7
// 0: 同一类型，有不同objAddress的信息元素集合 (一个地址+一个元素)
// 1： 同一类型，相同objAddress信息元素集合 (一个地址,后续连续N个元素)
type Variable byte

// Variable sequence
const (
	VariableSeq = 0x80
)

// Cause is the cause of transmission.
// See companion standard 101, subclause 7.2.3.
// Cause defined
// | T | P/N | 5..0 cause |
// T = test, 0: 未试验, 1：试验
// P/N 对由启动应用功能所请求的激活以肯定或者否定的确认 0: 肯定确认, 1: 否定确认
type Cause byte

// OriginAddr is originator address, See companion standard 101, subclause 7.2.3.
// The width is controlled by Params.CauseSize. width 2 includes/activates the originator address.
// <0>: 未用
// <1..255>: 源发地址
type OriginAddr byte

// The 2 most significant bits are flags.
const (
	// TestFlag marks the cause of transmission for testing. bit7
	TestFlag Cause = 0x80

	// NegFlag indicates the negative (or positive) confirmation
	// of activation requested by the primary application function. bit6
	NegFlag Cause = 0x40
)

// CauseOfTransmission is the cause of transmission.
type CauseOfTransmission struct {
	Cause      Cause
	IsTest     bool
	IsNegative bool
}

// Cause of transmission bit5-bit0
// <0> 未定义
// <1..63> 传输原因序号
// <1..47> 标准定义
// <48..63> 专用范围
// NOTE: 信息对象带或不带时标由标识符类型的不同序列来区别
const (
	Unused   Cause = iota // unused
	Percyc                // periodic, cyclic
	Back                  // background scan
	Spont                 // spontaneous
	Init                  // initialized
	Req                   // request or requested
	Act                   // activation
	Actcon                // activation confirmation
	Deact                 // deactivation
	Deactcon              // deactivation confirmation
	Actterm               // activation termination
	Retrem                // return information caused by a remote command
	Retloc                // return information caused by a local command
	File                  // file transfer
	Auth                  // authentication
	Seskey                // maintenance of authentication session key
	Usrkey                // maintenance of user role and update key
	_                     // reserved for further compatible definitions
	_                     // reserved for further compatible definitions
	_                     // reserved for further compatible definitions
	Inrogen               // interrogated by station interrogation
	Inro1                 // interrogated by group 1 interrogation
	Inro2                 // interrogated by group 2 interrogation
	Inro3                 // interrogated by group 3 interrogation
	Inro4                 // interrogated by group 4 interrogation
	Inro5                 // interrogated by group 5 interrogation
	Inro6                 // interrogated by group 6 interrogation
	Inro7                 // interrogated by group 7 interrogation
	Inro8                 // interrogated by group 8 interrogation
	Inro9                 // interrogated by group 9 interrogation
	Inro10                // interrogated by group 10 interrogation
	Inro11                // interrogated by group 11 interrogation
	Inro12                // interrogated by group 12 interrogation
	Inro13                // interrogated by group 13 interrogation
	Inro14                // interrogated by group 14 interrogation
	Inro15                // interrogated by group 15 interrogation
	Inro16                // interrogated by group 16 interrogation
	Reqcogen              // requested by general counter request
	Reqco1                // requested by group 1 counter request
	Reqco2                // requested by group 2 counter request
	Reqco3                // requested by group 3 counter request
	Reqco4                // requested by group 4 counter request
	_                     // reserved for further compatible definitions
	_                     // reserved for further compatible definitions
	UnkType               // unknown type identification
	UnkCause              // unknown cause of transmission
	UnkAddr               // unknown common address of ASDU
	UnkInfo               // unknown information object address
)

// Causal semantics description
var causeSemantics = []string{
	"unused0",
	"percyc",
	"back",
	"spont",
	"init",
	"req",
	"act",
	"actcon",
	"deact",
	"deactcon",
	"actterm",
	"retrem",
	"retloc",
	"file",
	"auth",
	"seskey",
	"usrkey",
	"reserved17",
	"reserved18",
	"reserved19",
	"inrogen",
	"inro1",
	"inro2",
	"inro3",
	"inro4",
	"inro5",
	"inro6",
	"inro7",
	"inro8",
	"inro9",
	"inro10",
	"inro11",
	"inro12",
	"inro13",
	"inro14",
	"inro15",
	"inro16",
	"reqcogen",
	"reqco1",
	"reqco2",
	"reqco3",
	"reqco4",
	"reserved42",
	"reserved43",
	"unktype",
	"unkcause",
	"unkaddr",
	"unkinfo",
	"special48",
	"special49",
	"special50",
	"special51",
	"special52",
	"special53",
	"special54",
	"special55",
	"special56",
	"special57",
	"special58",
	"special59",
	"special60",
	"special61",
	"special62",
	"special63",
}

// String 返回Cause的字符串,包含相应应用的",neg" and ",test"
func (this Cause) String() string {
	s := causeSemantics[this&^(NegFlag|TestFlag)]

	switch this & (NegFlag | TestFlag) {
	case NegFlag:
		s += ",neg"
	case TestFlag:
		s += ",test"
	case NegFlag | TestFlag:
		s += ",neg,test"
	}

	return s
}

// CommonAddr is a station address.
// The width is controlled by Params.CommonAddrSize.
// width 1
// <0>: 未用
// <1..254>: 站地址
// <255>: 全局地址
// width 2
// <0>: 未用
// <1..65534>: 站地址
// <65535>: 全局地址
type CommonAddr uint16

const (
	// InvalidCommonAddr is the invalid common address.
	InvalidCommonAddr CommonAddr = 0
	// GlobalCommonAddr is the broadcast address. Use is restricted
	// to C_IC_NA_1, C_CI_NA_1, C_CS_NA_1 and C_RP_NA_1.
	// When in 8-bit mode 255 is mapped to this value on the fly.
	GlobalCommonAddr CommonAddr = 65535
)
