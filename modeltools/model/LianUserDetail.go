package models


type LianUserDetail struct {
	Uid int `stbl:"uid" desc:"用户ID"` // 用户ID
	Phone int64 `stbl:"phone" desc:"手机号"` // 手机号
	Uname string `stbl:"uname" desc:"用户昵称"` // 用户昵称
	Sex int `stbl:"sex" desc:"性别: 1: 男 2: 女"` // 性别: 1: 男 2: 女
	Avatar string `stbl:"avatar" desc:"头像"` // 头像
	City string `stbl:"city" desc:"城市"` // 城市
	Intro string `stbl:"intro" desc:"自我介绍"` // 自我介绍
	Authorization int `stbl:"authorization" desc:"认证状态 0:财联社认证状态 1:蜂网认证状态 2:科创板APP"` // 认证状态 0:财联社认证状态 1:蜂网认证状态 2:科创板APP
	AttentionIndustry string `stbl:"attention_industry" desc:"关注行业"` // 关注行业
	Investment string `stbl:"investment" desc:"投入资本"` // 投入资本
	InvestStyle string `stbl:"invest_style" desc:"投资风格"` // 投资风格
	FengwangAuthorizeStatus int `stbl:"fengwang_authorize_status" desc:"蜂网认证状态 0: 默认 1: 认证中 2: 认证失败 3: 认证成功"` // 蜂网认证状态 0: 默认 1: 认证中 2: 认证失败 3: 认证成功
	AuthorizeStatus int `stbl:"authorize_status" desc:"财联社认证状态 0: 默认 1: 认证中 2: 认证失败 3: 认证成功"` // 财联社认证状态 0: 默认 1: 认证中 2: 认证失败 3: 认证成功
	CommentStatus int `stbl:"comment_status" desc:"0:正常用户 1:在黑名单中 2:广告用户 4:举报用户 二进制按位存储"` // 0:正常用户 1:在黑名单中 2:广告用户 4:举报用户 二进制按位存储
	Source int `stbl:"source" desc:"0财联社、1微博"` // 0财联社、1微博
	VipExpireTime int64 `stbl:"vip_expire_time" desc:"vip到期时间"` // vip到期时间
	Addressee sql.NullString `stbl:"addressee" desc:"收货地址"` // 收货地址
	CreditsNum sql.NullInt64 `stbl:"credits_num" desc:"积分数"` // 积分数
	WalletNum sql.NullInt64 `stbl:"wallet_num" desc:"钱包余额"` // 钱包余额
	Flag int `stbl:"flag" desc:"会员标记：0普通会员 1银卡 2金卡 3钻石卡"` // 会员标记：0普通会员 1银卡 2金卡 3钻石卡
	Birthday string `stbl:"birthday" desc:"生日"` // 生日
	Industry int `stbl:"industry" desc:"行业"` // 行业
	Career int `stbl:"career" desc:"职业"` // 职业
	MonthlySalary int `stbl:"monthly_salary" desc:"月收入"` // 月收入
	Education int `stbl:"education" desc:"学历"` // 学历
}

// uid, phone, uname, sex, avatar, city, intro, authorization, attention_industry, investment, invest_style, fengwang_authorize_status, authorize_status, comment_status, source, vip_expire_time, addressee, credits_num, wallet_num, flag, birthday, industry, career, monthly_salary, education, 
// uid=?, phone=?, uname=?, sex=?, avatar=?, city=?, intro=?, authorization=?, attention_industry=?, investment=?, invest_style=?, fengwang_authorize_status=?, authorize_status=?, comment_status=?, source=?, vip_expire_time=?, addressee=?, credits_num=?, wallet_num=?, flag=?, birthday=?, industry=?, career=?, monthly_salary=?, education=?, 
// a.uid, a.phone, a.uname, a.sex, a.avatar, a.city, a.intro, a.authorization, a.attention_industry, a.investment, a.invest_style, a.fengwang_authorize_status, a.authorize_status, a.comment_status, a.source, a.vip_expire_time, a.addressee, a.credits_num, a.wallet_num, a.flag, a.birthday, a.industry, a.career, a.monthly_salary, a.education, 

 p.Uid = req.Uid
 p.Phone = req.Phone
 p.Uname = req.Uname
 p.Sex = req.Sex
 p.Avatar = req.Avatar
 p.City = req.City
 p.Intro = req.Intro
 p.Authorization = req.Authorization
 p.AttentionIndustry = req.AttentionIndustry
 p.Investment = req.Investment
 p.InvestStyle = req.InvestStyle
 p.FengwangAuthorizeStatus = req.FengwangAuthorizeStatus
 p.AuthorizeStatus = req.AuthorizeStatus
 p.CommentStatus = req.CommentStatus
 p.Source = req.Source
 p.VipExpireTime = req.VipExpireTime
 p.Addressee = req.Addressee
 p.CreditsNum = req.CreditsNum
 p.WalletNum = req.WalletNum
 p.Flag = req.Flag
 p.Birthday = req.Birthday
 p.Industry = req.Industry
 p.Career = req.Career
 p.MonthlySalary = req.MonthlySalary
 p.Education = req.Education
 
// &vo.Uid, &vo.Phone, &vo.Uname, &vo.Sex, &vo.Avatar, &vo.City, &vo.Intro, &vo.Authorization, &vo.AttentionIndustry, &vo.Investment, &vo.InvestStyle, &vo.FengwangAuthorizeStatus, &vo.AuthorizeStatus, &vo.CommentStatus, &vo.Source, &vo.VipExpireTime, &vo.Addressee, &vo.CreditsNum, &vo.WalletNum, &vo.Flag, &vo.Birthday, &vo.Industry, &vo.Career, &vo.MonthlySalary, &vo.Education, 

// @Param uid formData int true "用户ID"
// @Param phone formData int64 true "手机号"
// @Param uname formData string true "用户昵称"
// @Param sex formData int true "性别: 1: 男 2: 女"
// @Param avatar formData string true "头像"
// @Param city formData string true "城市"
// @Param intro formData string true "自我介绍"
// @Param authorization formData int true "认证状态 0:财联社认证状态 1:蜂网认证状态 2:科创板APP"
// @Param attention_industry formData string true "关注行业"
// @Param investment formData string true "投入资本"
// @Param invest_style formData string true "投资风格"
// @Param fengwang_authorize_status formData int true "蜂网认证状态 0: 默认 1: 认证中 2: 认证失败 3: 认证成功"
// @Param authorize_status formData int true "财联社认证状态 0: 默认 1: 认证中 2: 认证失败 3: 认证成功"
// @Param comment_status formData int true "0:正常用户 1:在黑名单中 2:广告用户 4:举报用户 二进制按位存储"
// @Param source formData int true "0财联社、1微博"
// @Param vip_expire_time formData int64 true "vip到期时间"
// @Param addressee formData string true "收货地址"
// @Param credits_num formData int true "积分数"
// @Param wallet_num formData int true "钱包余额"
// @Param flag formData int true "会员标记：0普通会员 1银卡 2金卡 3钻石卡"
// @Param birthday formData string true "生日"
// @Param industry formData int true "行业"
// @Param career formData int true "职业"
// @Param monthly_salary formData int true "月收入"
// @Param education formData int true "学历"

type LianUserDetail struct {
	Uid int `json:"uid" desc:"用户ID"` // 用户ID
	Phone int64 `json:"phone" desc:"手机号"` // 手机号
	Uname string `json:"uname" desc:"用户昵称"` // 用户昵称
	Sex int `json:"sex" desc:"性别: 1: 男 2: 女"` // 性别: 1: 男 2: 女
	Avatar string `json:"avatar" desc:"头像"` // 头像
	City string `json:"city" desc:"城市"` // 城市
	Intro string `json:"intro" desc:"自我介绍"` // 自我介绍
	Authorization int `json:"authorization" desc:"认证状态 0:财联社认证状态 1:蜂网认证状态 2:科创板APP"` // 认证状态 0:财联社认证状态 1:蜂网认证状态 2:科创板APP
	AttentionIndustry string `json:"attention_industry" desc:"关注行业"` // 关注行业
	Investment string `json:"investment" desc:"投入资本"` // 投入资本
	InvestStyle string `json:"invest_style" desc:"投资风格"` // 投资风格
	FengwangAuthorizeStatus int `json:"fengwang_authorize_status" desc:"蜂网认证状态 0: 默认 1: 认证中 2: 认证失败 3: 认证成功"` // 蜂网认证状态 0: 默认 1: 认证中 2: 认证失败 3: 认证成功
	AuthorizeStatus int `json:"authorize_status" desc:"财联社认证状态 0: 默认 1: 认证中 2: 认证失败 3: 认证成功"` // 财联社认证状态 0: 默认 1: 认证中 2: 认证失败 3: 认证成功
	CommentStatus int `json:"comment_status" desc:"0:正常用户 1:在黑名单中 2:广告用户 4:举报用户 二进制按位存储"` // 0:正常用户 1:在黑名单中 2:广告用户 4:举报用户 二进制按位存储
	Source int `json:"source" desc:"0财联社、1微博"` // 0财联社、1微博
	VipExpireTime int64 `json:"vip_expire_time" desc:"vip到期时间"` // vip到期时间
	Addressee string `json:"addressee" desc:"收货地址"` // 收货地址
	CreditsNum int `json:"credits_num" desc:"积分数"` // 积分数
	WalletNum int `json:"wallet_num" desc:"钱包余额"` // 钱包余额
	Flag int `json:"flag" desc:"会员标记：0普通会员 1银卡 2金卡 3钻石卡"` // 会员标记：0普通会员 1银卡 2金卡 3钻石卡
	Birthday string `json:"birthday" desc:"生日"` // 生日
	Industry int `json:"industry" desc:"行业"` // 行业
	Career int `json:"career" desc:"职业"` // 职业
	MonthlySalary int `json:"monthly_salary" desc:"月收入"` // 月收入
	Education int `json:"education" desc:"学历"` // 学历
}
