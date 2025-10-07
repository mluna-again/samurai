package main

import _ "embed"

// i know this looks terrible but surely its faster than reading an embed.FS at runtime, no? too lazy to test.

//go:embed ascii/samurai-lg1.ascii
var samLg1 string

//go:embed ascii/samurai-lg2.ascii
var samLg2 string

//go:embed ascii/samurai-lg3.ascii
var samLg3 string

//go:embed ascii/samurai-lg4.ascii
var samLg4 string

//go:embed ascii/samurai-lg5.ascii
var samLg5 string

//go:embed ascii/samurai-lg6.ascii
var samLg6 string

//go:embed ascii/samurai-lg7.ascii
var samLg7 string

//go:embed ascii/samurai-lg8.ascii
var samLg8 string

//go:embed ascii/samurai-lg9.ascii
var samLg9 string

//go:embed ascii/samurai-lg10.ascii
var samLg10 string

//go:embed ascii/samurai-lg11.ascii
var samLg11 string

//go:embed ascii/samurai-lg12.ascii
var samLg12 string

//go:embed ascii/samurai-lg13.ascii
var samLg13 string

//go:embed ascii/samurai-lg14.ascii
var samLg14 string

//go:embed ascii/samurai-lg15.ascii
var samLg15 string

//go:embed ascii/samurai-lg16.ascii
var samLg16 string

//go:embed ascii/samurai-lg17.ascii
var samLg17 string

//go:embed ascii/samurai-lg18.ascii
var samLg18 string

//go:embed ascii/samurai-lg19.ascii
var samLg19 string

//go:embed ascii/samurai-lg20.ascii
var samLg20 string

//go:embed ascii/samurai-lg21.ascii
var samLg21 string

//go:embed ascii/samurai-lg22.ascii
var samLg22 string

//go:embed ascii/samurai-lg23.ascii
var samLg23 string

//go:embed ascii/samurai-lg24.ascii
var samLg24 string

//go:embed ascii/samurai-lg25.ascii
var samLg25 string

//go:embed ascii/samurai-md1.ascii
var samMd1 string

//go:embed ascii/samurai-md2.ascii
var samMd2 string

//go:embed ascii/samurai-md3.ascii
var samMd3 string

//go:embed ascii/samurai-md4.ascii
var samMd4 string

//go:embed ascii/samurai-md5.ascii
var samMd5 string

//go:embed ascii/samurai-md6.ascii
var samMd6 string

//go:embed ascii/samurai-md7.ascii
var samMd7 string

//go:embed ascii/samurai-md8.ascii
var samMd8 string

//go:embed ascii/samurai-md9.ascii
var samMd9 string

//go:embed ascii/samurai-md10.ascii
var samMd10 string

//go:embed ascii/samurai-md11.ascii
var samMd11 string

//go:embed ascii/samurai-md12.ascii
var samMd12 string

//go:embed ascii/samurai-md13.ascii
var samMd13 string

//go:embed ascii/samurai-md14.ascii
var samMd14 string

//go:embed ascii/samurai-md15.ascii
var samMd15 string

//go:embed ascii/samurai-md16.ascii
var samMd16 string

//go:embed ascii/samurai-md17.ascii
var samMd17 string

//go:embed ascii/samurai-md18.ascii
var samMd18 string

//go:embed ascii/samurai-md19.ascii
var samMd19 string

//go:embed ascii/samurai-md20.ascii
var samMd20 string

//go:embed ascii/samurai-md21.ascii
var samMd21 string

//go:embed ascii/samurai-md22.ascii
var samMd22 string

//go:embed ascii/samurai-md23.ascii
var samMd23 string

//go:embed ascii/samurai-md24.ascii
var samMd24 string

//go:embed ascii/samurai-md25.ascii
var samMd25 string

//go:embed ascii/samurai-sm1.ascii
var samSm1 string

//go:embed ascii/samurai-sm2.ascii
var samSm2 string

//go:embed ascii/samurai-sm3.ascii
var samSm3 string

//go:embed ascii/samurai-sm4.ascii
var samSm4 string

//go:embed ascii/samurai-sm5.ascii
var samSm5 string

//go:embed ascii/samurai-sm6.ascii
var samSm6 string

//go:embed ascii/samurai-sm7.ascii
var samSm7 string

//go:embed ascii/samurai-sm8.ascii
var samSm8 string

//go:embed ascii/samurai-sm9.ascii
var samSm9 string

//go:embed ascii/samurai-sm10.ascii
var samSm10 string

//go:embed ascii/samurai-sm11.ascii
var samSm11 string

//go:embed ascii/samurai-sm12.ascii
var samSm12 string

//go:embed ascii/samurai-sm13.ascii
var samSm13 string

//go:embed ascii/samurai-sm14.ascii
var samSm14 string

//go:embed ascii/samurai-sm15.ascii
var samSm15 string

//go:embed ascii/samurai-sm16.ascii
var samSm16 string

//go:embed ascii/samurai-sm17.ascii
var samSm17 string

//go:embed ascii/samurai-sm18.ascii
var samSm18 string

//go:embed ascii/samurai-sm19.ascii
var samSm19 string

//go:embed ascii/samurai-sm20.ascii
var samSm20 string

//go:embed ascii/samurai-sm21.ascii
var samSm21 string

//go:embed ascii/samurai-sm22.ascii
var samSm22 string

//go:embed ascii/samurai-sm23.ascii
var samSm23 string

//go:embed ascii/samurai-sm24.ascii
var samSm24 string

//go:embed ascii/samurai-sm25.ascii
var samSm25 string

var samuraiLarge = []string{
	samLg1,
	samLg2,
	samLg3,
	samLg4,
	samLg5,
	samLg6,
	samLg7,
	samLg8,
	samLg9,
	samLg10,
	samLg11,
	samLg12,
	samLg13,
	samLg14,
	samLg15,
	samLg16,
	samLg17,
	samLg18,
	samLg19,
	samLg20,
	samLg21,
	samLg22,
	samLg23,
	samLg24,
	samLg25,
}

var samuraiMedium = []string{
	samMd1,
	samMd2,
	samMd3,
	samMd4,
	samMd5,
	samMd6,
	samMd7,
	samMd8,
	samMd9,
	samMd10,
	samMd11,
	samMd12,
	samMd13,
	samMd14,
	samMd15,
	samMd16,
	samMd17,
	samMd18,
	samMd19,
	samMd20,
	samMd21,
	samMd22,
	samMd23,
	samMd24,
	samMd25,
}

var samuraiSmall = []string{
	samSm1,
	samSm2,
	samSm3,
	samSm4,
	samSm5,
	samSm6,
	samSm7,
	samSm8,
	samSm9,
	samSm10,
	samSm11,
	samSm12,
	samSm13,
	samSm14,
	samSm15,
	samSm16,
	samSm17,
	samSm18,
	samSm19,
	samSm20,
	samSm21,
	samSm22,
	samSm23,
	samSm24,
	samSm25,
}
