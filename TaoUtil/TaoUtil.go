// @Title TaoUtil
// @Description 道历工具
// @Author 6tail
package TaoUtil

// 三会日
var SAN_HUI = []string{"1-7", "7-7", "10-15"}

// 三元日
var SAN_YUAN = []string{"1-15", "7-15", "10-15"}

// 五腊日
var WU_LA = []string{"1-1", "5-5", "7-7", "10-1", "12-8"}

// 暗戊
var AN_WU = []string{"未", "戌", "辰", "寅", "午", "子", "酉", "申", "巳", "亥", "卯", "丑"}

// 八会日
var BA_HUI = map[string]string{
	"丙午": "天会",
	"壬午": "地会",
	"壬子": "人会",
	"庚午": "日会",
	"庚申": "月会",
	"辛酉": "星辰会",
	"甲辰": "五行会",
	"甲戌": "四时会",
}

// 八节日
var BA_JIE = map[string]string{
	"立春": "东北方度仙上圣天尊同梵炁始青天君下降",
	"春分": "东方玉宝星上天尊同青帝九炁天君下降",
	"立夏": "东南方好生度命天尊同梵炁始丹天君下降",
	"夏至": "南方玄真万福天尊同赤帝三炁天君下降",
	"立秋": "西南方太灵虚皇天尊同梵炁始素天君下降",
	"秋分": "西方太妙至极天尊同白帝七炁天君下降",
	"立冬": "西北方无量太华天尊同梵炁始玄天君下降",
	"冬至": "北方玄上玉宸天尊同黑帝五炁天君下降",
}

// 节日
var FESTIVAL = map[string][][]string{
	"1-1":   {{"天腊之辰", "天腊，此日五帝会于束方九炁青天"}},
	"1-3":   {{"郝真人圣诞"}, {"孙真人圣诞"}},
	"1-5":   {{"孙祖清静元君诞"}},
	"1-7":   {{"举迁赏会", "此日上元赐福，天官同地水二官考校罪福"}},
	"1-9":   {{"玉皇上帝圣诞"}},
	"1-13":  {{"关圣帝君飞升"}},
	"1-15":  {{"上元天官圣诞"}, {"老祖天师圣诞"}},
	"1-19":  {{"长春邱真人(邱处机}圣诞"}},
	"1-28":  {{"许真君(许逊天师}圣诞"}},
	"2-1":   {{"勾陈天皇大帝圣诞"}, {"长春刘真人(刘渊然}圣诞"}},
	"2-2":   {{"土地正神诞"}, {"姜太公圣诞"}},
	"2-3":   {{"文昌梓潼帝君圣诞"}},
	"2-6":   {{"东华帝君圣诞"}},
	"2-13":  {{"度人无量葛真君圣诞"}},
	"2-15":  {{"太清道德天尊(太上老君}圣诞"}},
	"2-19":  {{"慈航真人圣诞"}},
	"3-1":   {{"谭祖(谭处端}长真真人圣诞"}},
	"3-3":   {{"玄天上帝圣诞"}},
	"3-6":   {{"眼光娘娘圣诞"}},
	"3-15":  {{"天师张大真人圣诞"}, {"财神赵公元帅圣诞"}},
	"3-16":  {{"三茅真君得道之辰"}, {"中岳大帝圣诞"}},
	"3-18":  {{"王祖(王处一}玉阳真人圣诞"}, {"后土娘娘圣诞"}},
	"3-19":  {{"太阳星君圣诞"}},
	"3-20":  {{"子孙娘娘圣诞"}},
	"3-23":  {{"天后妈祖圣诞"}},
	"3-26":  {{"鬼谷先师诞"}},
	"3-28":  {{"东岳大帝圣诞"}},
	"4-1":   {{"长生谭真君成道之辰"}},
	"4-10":  {{"何仙姑圣诞"}},
	"4-14":  {{"吕祖纯阳祖师圣诞"}},
	"4-15":  {{"钟离祖师圣诞"}},
	"4-18":  {{"北极紫微大帝圣诞"}, {"泰山圣母碧霞元君诞"}, {"华佗神医先师诞"}},
	"4-20":  {{"眼光圣母娘娘诞"}},
	"4-28":  {{"神农先帝诞"}},
	"5-1":   {{"南极长生大帝圣诞"}},
	"5-5":   {{"地腊之辰", "地腊，此日五帝会於南方三炁丹天"}, {"南方雷祖圣诞"}, {"地祗温元帅圣诞"}, {"雷霆邓天君圣诞"}},
	"5-11":  {{"城隍爷圣诞"}},
	"5-13":  {{"关圣帝君降神"}, {"关平太子圣诞"}},
	"5-18":  {{"张天师圣诞"}},
	"5-20":  {{"马祖丹阳真人圣诞"}},
	"5-29":  {{"紫青白祖师圣诞"}},
	"6-1":   {{"南斗星君下降"}},
	"6-2":   {{"南斗星君下降"}},
	"6-3":   {{"南斗星君下降"}},
	"6-4":   {{"南斗星君下降"}},
	"6-5":   {{"南斗星君下降"}},
	"6-6":   {{"南斗星君下降"}},
	"6-10":  {{"刘海蟾祖师圣诞"}},
	"6-15":  {{"灵官王天君圣诞"}},
	"6-19":  {{"慈航(观音}成道日"}},
	"6-23":  {{"火神圣诞"}},
	"6-24":  {{"南极大帝中方雷祖圣诞"}, {"关圣帝君圣诞"}},
	"6-26":  {{"二郎真君圣诞"}},
	"7-7":   {{"道德腊之辰", "道德腊，此日五帝会于西方七炁素天"}, {"庆生中会", "此日中元赦罪，地官同天水二官考校罪福"}},
	"7-12":  {{"西方雷祖圣诞"}},
	"7-15":  {{"中元地官大帝圣诞"}},
	"7-18":  {{"王母娘娘圣诞"}},
	"7-20":  {{"刘祖(刘处玄}长生真人圣诞"}},
	"7-22":  {{"财帛星君文财神增福相公李诡祖圣诞"}},
	"7-26":  {{"张三丰祖师圣诞"}},
	"8-1":   {{"许真君飞升日"}},
	"8-3":   {{"九天司命灶君诞"}},
	"8-5":   {{"北方雷祖圣诞"}},
	"8-10":  {{"北岳大帝诞辰"}},
	"8-15":  {{"太阴星君诞"}},
	"9-1":   {{"北斗九皇降世之辰"}},
	"9-2":   {{"北斗九皇降世之辰"}},
	"9-3":   {{"北斗九皇降世之辰"}},
	"9-4":   {{"北斗九皇降世之辰"}},
	"9-5":   {{"北斗九皇降世之辰"}},
	"9-6":   {{"北斗九皇降世之辰"}},
	"9-7":   {{"北斗九皇降世之辰"}},
	"9-8":   {{"北斗九皇降世之辰"}},
	"9-9":   {{"北斗九皇降世之辰"}, {"斗姥元君圣诞"}, {"重阳帝君圣诞"}, {"玄天上帝飞升"}, {"酆都大帝圣诞"}},
	"9-22":  {{"增福财神诞"}},
	"9-23":  {{"萨翁真君圣诞"}},
	"9-28":  {{"五显灵官马元帅圣诞"}},
	"10-1":  {{"民岁腊之辰", "民岁腊，此日五帝会於北方五炁黑天"}, {"东皇大帝圣诞"}},
	"10-3":  {{"三茅应化真君圣诞"}},
	"10-6":  {{"天曹诸司五岳五帝圣诞"}},
	"10-15": {{"下元水官大帝圣诞"}, {"建生大会", "此日下元解厄，水官同天地二官考校罪福"}},
	"10-18": {{"地母娘娘圣诞"}},
	"10-19": {{"长春邱真君飞升"}},
	"10-20": {{"虚靖天师(即三十代天师弘悟张真人}诞"}},
	"11-6":  {{"西岳大帝圣诞"}},
	"11-9":  {{"湘子韩祖圣诞"}},
	"11-11": {{"太乙救苦天尊圣诞"}},
	"11-26": {{"北方五道圣诞"}},
	"12-8":  {{"王侯腊之辰", "王侯腊，此日五帝会於上方玄都玉京"}},
	"12-16": {{"南岳大帝圣诞"}, {"福德正神诞"}},
	"12-20": {{"鲁班先师圣诞"}},
	"12-21": {{"天猷上帝圣诞"}},
	"12-22": {{"重阳祖师圣诞"}},
	"12-23": {{"祭灶王", "最适宜谢旧年太岁，开启拜新年太岁"}},
	"12-25": {{"玉帝巡天"}, {"天神下降"}},
	"12-29": {{"清静孙真君(孙不二}成道"}},
}
