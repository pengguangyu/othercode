package models


type LianTopnewsFeedRecommend struct {
	Id int `stbl:"id" desc:""` 
	Pid int `stbl:"pid" desc:"群组标记，0默认， 非0表示此条隶属于pid指向的记录，比如 盯盘 盯热点"` // 群组标记，0默认， 非0表示此条隶属于pid指向的记录，比如 盯盘 盯热点
	Type int `stbl:"type" desc:"类型，1VIP文章，2电报有图，3电报无图，4专题，5要问，6话题文章，7横向直播/视频，8资讯文章，9盯热点， 10热门话题，11二十四小时热榜"` // 类型，1VIP文章，2电报有图，3电报无图，4专题，5要问，6话题文章，7横向直播/视频，8资讯文章，9盯热点， 10热门话题，11二十四小时热榜
	Sort int `stbl:"sort" desc:"展示顺序"` // 展示顺序
	Cid int `stbl:"cid" desc:"内容id"` // 内容id
	Cid2 int `stbl:"cid2" desc:"内容id2"` // 内容id2
	Word string `stbl:"word" desc:"类型内容展示名称（默认文章标题，支持修改）"` // 类型内容展示名称（默认文章标题，支持修改）
	Cover string `stbl:"cover" desc:"封面"` // 封面
	Del int `stbl:"del" desc:"软删除 0正常 1删除"` // 软删除 0正常 1删除
	CreateTime time.Time `stbl:"create_time" desc:"创建时间"` // 创建时间
	ModifyedTime time.Time `stbl:"modifyed_time" desc:"更新时间"` // 更新时间
}

// id, pid, type, sort, cid, cid2, word, cover, del, create_time, modifyed_time, 
// id=?, pid=?, type=?, sort=?, cid=?, cid2=?, word=?, cover=?, del=?, create_time=?, modifyed_time=?, 
// a.id, a.pid, a.type, a.sort, a.cid, a.cid2, a.word, a.cover, a.del, a.create_time, a.modifyed_time, 

 p.Id = req.Id
 p.Pid = req.Pid
 p.Type = req.Type
 p.Sort = req.Sort
 p.Cid = req.Cid
 p.Cid2 = req.Cid2
 p.Word = req.Word
 p.Cover = req.Cover
 p.Del = req.Del
 p.CreateTime = req.CreateTime
 p.ModifyedTime = req.ModifyedTime
 
// &vo.Id, &vo.Pid, &vo.Type, &vo.Sort, &vo.Cid, &vo.Cid2, &vo.Word, &vo.Cover, &vo.Del, &vo.CreateTime, &vo.ModifyedTime, 

// @Param id formData int true ""
// @Param pid formData int true "群组标记，0默认， 非0表示此条隶属于pid指向的记录，比如 盯盘 盯热点"
// @Param type formData int true "类型，1VIP文章，2电报有图，3电报无图，4专题，5要问，6话题文章，7横向直播/视频，8资讯文章，9盯热点， 10热门话题，11二十四小时热榜"
// @Param sort formData int true "展示顺序"
// @Param cid formData int true "内容id"
// @Param cid2 formData int true "内容id2"
// @Param word formData string true "类型内容展示名称（默认文章标题，支持修改）"
// @Param cover formData string true "封面"
// @Param del formData int true "软删除 0正常 1删除"
// @Param create_time formData int64 true "创建时间"
// @Param modifyed_time formData int64 true "更新时间"

type LianTopnewsFeedRecommend struct {
	Id int `json:"id" desc:""` 
	Pid int `json:"pid" desc:"群组标记，0默认， 非0表示此条隶属于pid指向的记录，比如 盯盘 盯热点"` // 群组标记，0默认， 非0表示此条隶属于pid指向的记录，比如 盯盘 盯热点
	Type int `json:"type" desc:"类型，1VIP文章，2电报有图，3电报无图，4专题，5要问，6话题文章，7横向直播/视频，8资讯文章，9盯热点， 10热门话题，11二十四小时热榜"` // 类型，1VIP文章，2电报有图，3电报无图，4专题，5要问，6话题文章，7横向直播/视频，8资讯文章，9盯热点， 10热门话题，11二十四小时热榜
	Sort int `json:"sort" desc:"展示顺序"` // 展示顺序
	Cid int `json:"cid" desc:"内容id"` // 内容id
	Cid2 int `json:"cid2" desc:"内容id2"` // 内容id2
	Word string `json:"word" desc:"类型内容展示名称（默认文章标题，支持修改）"` // 类型内容展示名称（默认文章标题，支持修改）
	Cover string `json:"cover" desc:"封面"` // 封面
	Del int `json:"del" desc:"软删除 0正常 1删除"` // 软删除 0正常 1删除
	CreateTime int64 `json:"create_time" desc:"创建时间"` // 创建时间
	ModifyedTime int64 `json:"modifyed_time" desc:"更新时间"` // 更新时间
}
