package main

var symbolTemplates = map[string]Symbols{
	"compatible": {
		Lock:                 "RO",
		Network:              "SSH",
		Separator:            "\u25B6",
		SeparatorThin:        "\u276F",
		SeparatorReverse:     "\u25C0",
		SeparatorReverseThin: "\u276E",

		RepoDetached:   "\u2693",
		RepoAhead:      "\u2B06",
		RepoBehind:     "\u2B07",
		RepoStaged:     "\u2714",
		RepoNotStaged:  "\u270E",
		RepoUntracked:  "+",
		RepoConflicted: "\u273C",
		RepoStashed:    "\u2691",
	},
	"patched": {
		Lock:                 "\uE0A2",
		Network:              "\uE0A2",
		Separator:            "\uE0B0",
		SeparatorThin:        "\uE0B1",
		SeparatorReverse:     "\uE0B2",
		SeparatorReverseThin: "\uE0B3",

		RepoDetached:   "\u2693",
		RepoAhead:      "\u2B06",
		RepoBehind:     "\u2B07",
		RepoStaged:     "\u2714",
		RepoNotStaged:  "\u270E",
		RepoUntracked:  "+",
		RepoConflicted: "\u273C",
		RepoStashed:    "\u2691",
	},
	"flat": {
		RepoDetached:   "\u2693",
		RepoAhead:      "\u2B06",
		RepoBehind:     "\u2B07",
		RepoStaged:     "\u2714",
		RepoNotStaged:  "\u270E",
		RepoUntracked:  "+",
		RepoConflicted: "\u273C",
		RepoStashed:    "\u2691",
	},
}

var shellInfos = map[string]ShellInfo{
	"bash": {
		colorTemplate:    "\\[\\e%s\\]",
		rootIndicator:    "\\$",
		escapedBackslash: `\\\\`,
		escapedBacktick:  "\\`",
		escapedDollar:    `\$`,
		evalPromptPrefix: `PS1="`,
		evalPromptSuffix: `"`,
	},
	"zsh": {
		colorTemplate:         "%%{\u001b%s%%}",
		rootIndicator:         "%#",
		escapedBackslash:      `\\`,
		escapedBacktick:       "\\`",
		escapedDollar:         `\$`,
		evalPromptPrefix:      `PROMPT="`,
		evalPromptSuffix:      `"`,
		evalPromptRightPrefix: `RPROMPT="`,
		evalPromptRightSuffix: `"`,
	},
	"bare": {
		colorTemplate: "%s",
		rootIndicator:    "$",
		escapedBackslash: `\`,
		escapedBacktick:  "`",
		escapedDollar:    `$`,
	},
}

var themes = map[string]Theme{
	"default": {
		Reset: 0xFF,

		UsernameFg:     250,
		UsernameBg:     0,
		UsernameRootBg: 124,

		HostnameFg: 250,
		HostnameBg: 238,

		HomeSpecialDisplay: true,
		HomeFg:             15,  // white
		HomeBg:             4,  // blueish
		AliasFg:            15,  // white
		AliasBg:            31,  // blueish
		PathFg:             250, // light grey
		PathBg:             237, // dark grey
		CwdFg:              254, // nearly-white grey
		SeparatorFg:        244,

		ReadonlyFg: 254,
		ReadonlyBg: 124,

		SshFg: 254,
		SshBg: 166, // medium orange

		DockerMachineFg: 177, // light purple
		DockerMachineBg: 55,  // purple

		KubeClusterFg:   117,
		KubeClusterBg:   26,
		KubeNamespaceFg: 170,
		KubeNamespaceBg: 17,

		DotEnvFg: 15, // white
		DotEnvBg: 55, // purple

		AWSFg: 15, // white
		AWSBg: 22, // dark green

		RepoCleanFg: 0,   // black
		RepoCleanBg: 148, // a light green color
		RepoDirtyFg: 15,  // white
		RepoDirtyBg: 161, // pink/red

		JobsFg: 39,
		JobsBg: 238,

		CmdPassedFg: 15,
		CmdPassedBg: 70,
		CmdFailedFg: 15,
		CmdFailedBg: 9,

		SvnChangesFg: 22, // dark green
		SvnChangesBg: 148,

		GitAheadFg:      250,
		GitAheadBg:      240,
		GitBehindFg:     250,
		GitBehindBg:     240,
		GitStagedFg:     15,
		GitStagedBg:     22,
		GitNotStagedFg:  15,
		GitNotStagedBg:  130,
		GitUntrackedFg:  15,
		GitUntrackedBg:  52,
		GitConflictedFg: 15,
		GitConflictedBg: 9,
		GitStashedFg:    15,
		GitStashedBg:    20,

		VirtualEnvFg: 00,
		VirtualEnvBg: 35, // a mid-tone green

		VirtualGoFg: 00,
		VirtualGoBg: 35,

		PerlbrewFg: 00,
		PerlbrewBg: 20, // a mid-tone blue

		TFWsFg: 15, // white
		TFWsBg: 26, // blue

		TimeFg: 15,
		TimeBg: 236,

		ShellVarFg: 52,
		ShellVarBg: 11,

		NodeFg: 15,
		NodeBg: 40,

		LoadFg:           15,
		LoadBg:           22,
		LoadHighBg:       161,
		LoadAvgValue:     5,
		LoadThresholdBad: 1.0,

		NixShellFg: 00,
		NixShellBg: 69, // a light blue

		DurationFg: 250,
		DurationBg: 237,

		HostnameColorizedFgMap: map[uint8]uint8{
			0:   250,
			1:   250,
			2:   120,
			3:   228,
			4:   250,
			5:   250,
			6:   123,
			7:   238,
			8:   0,
			9:   0,
			10:  0,
			11:  0,
			12:  250,
			13:  0,
			14:  0,
			15:  242,
			16:  250,
			17:  250,
			18:  250,
			19:  189,
			20:  254,
			21:  250,
			22:  83,
			23:  87,
			24:  117,
			25:  188,
			26:  254,
			27:  0,
			28:  120,
			29:  122,
			30:  123,
			31:  159,
			32:  255,
			33:  0,
			34:  157,
			35:  158,
			36:  159,
			37:  159,
			38:  195,
			39:  0,
			40:  194,
			41:  194,
			42:  195,
			43:  195,
			44:  195,
			45:  0,
			46:  0,
			47:  0,
			48:  0,
			49:  0,
			50:  0,
			51:  0,
			52:  250,
			53:  250,
			54:  250,
			55:  189,
			56:  254,
			57:  250,
			58:  227,
			59:  253,
			60:  255,
			61:  0,
			62:  233,
			63:  17,
			64:  192,
			65:  255,
			66:  195,
			67:  232,
			68:  233,
			69:  17,
			70:  193,
			71:  232,
			72:  232,
			73:  232,
			74:  234,
			75:  236,
			76:  194,
			77:  235,
			78:  235,
			79:  235,
			80:  235,
			81:  237,
			82:  0,
			83:  237,
			84:  237,
			85:  237,
			86:  237,
			87:  237,
			88:  250,
			89:  250,
			90:  250,
			91:  189,
			92:  254,
			93:  0,
			94:  222,
			95:  255,
			96:  255,
			97:  232,
			98:  233,
			99:  17,
			100: 228,
			101: 15,
			102: 232,
			103: 233,
			104: 17,
			105: 18,
			106: 229,
			107: 232,
			108: 234,
			109: 234,
			110: 236,
			111: 54,
			112: 230,
			113: 235,
			114: 22,
			115: 237,
			116: 238,
			117: 238,
			118: 0,
			119: 237,
			120: 22,
			121: 23,
			122: 23,
			123: 23,
			124: 252,
			125: 252,
			126: 189,
			127: 189,
			128: 254,
			129: 0,
			130: 223,
			131: 232,
			132: 232,
			133: 232,
			134: 233,
			135: 17,
			136: 229,
			137: 232,
			138: 233,
			139: 234,
			140: 53,
			141: 18,
			142: 229,
			143: 232,
			144: 234,
			145: 236,
			146: 17,
			147: 19,
			148: 230,
			149: 235,
			150: 238,
			151: 22,
			152: 23,
			153: 24,
			154: 0,
			155: 237,
			156: 22,
			157: 2,
			158: 29,
			159: 6,
			160: 254,
			161: 254,
			162: 254,
			163: 254,
			164: 254,
			165: 0,
			166: 255,
			167: 233,
			168: 233,
			169: 234,
			170: 234,
			171: 235,
			172: 230,
			173: 234,
			174: 52,
			175: 235,
			176: 53,
			177: 53,
			178: 230,
			179: 235,
			180: 236,
			181: 52,
			182: 53,
			183: 55,
			184: 230,
			185: 235,
			186: 238,
			187: 58,
			188: 240,
			189: 20,
			190: 0,
			191: 238,
			192: 58,
			193: 64,
			194: 35,
			195: 66,
			196: 0,
			197: 0,
			198: 0,
			199: 0,
			200: 0,
			201: 0,
			202: 0,
			203: 235,
			204: 235,
			205: 235,
			206: 235,
			207: 53,
			208: 0,
			209: 236,
			210: 52,
			211: 237,
			212: 53,
			213: 53,
			214: 0,
			215: 236,
			216: 238,
			217: 1,
			218: 89,
			219: 5,
			220: 0,
			221: 237,
			222: 58,
			223: 95,
			224: 131,
			225: 126,
			226: 0,
			227: 238,
			228: 58,
			229: 3,
			230: 143,
			231: 242,
			232: 250,
			233: 250,
			234: 250,
			235: 250,
			236: 250,
			237: 250,
			238: 251,
			239: 252,
			240: 188,
			241: 254,
			242: 254,
			243: 255,
			244: 0,
			245: 232,
			246: 233,
			247: 234,
			248: 235,
			249: 236,
			250: 237,
			251: 238,
			252: 239,
			253: 240,
			254: 242,
			255: 243,
		},
	},
	"low-contrast": {
		Reset: 0xFF,

		UsernameFg:     234,
		UsernameBg:     250,
		UsernameRootBg: 198,

		HostnameFg: 234,
		HostnameBg: 252,

		HomeSpecialDisplay: true,
		HomeFg:             30, // blueish-green
		HomeBg:             15, // white
		AliasFg:            30, // blueish-green
		AliasBg:            15, // white
		PathFg:             234,
		PathBg:             254,
		CwdFg:              236,
		SeparatorFg:        244,

		ReadonlyFg: 124,
		ReadonlyBg: 253,

		SshFg: 166, // medium orange
		SshBg: 254,

		DockerMachineFg: 55,  // purple
		DockerMachineBg: 177, // light purple

		KubeClusterFg:   117,
		KubeClusterBg:   26,
		KubeNamespaceFg: 170,
		KubeNamespaceBg: 17,

		DotEnvFg: 15, // white
		DotEnvBg: 55, // purple

		RepoCleanFg: 232, // black
		RepoCleanBg: 230, // light yellow
		RepoDirtyFg: 232, // black
		RepoDirtyBg: 223, // orange/peach

		JobsFg: 238,
		JobsBg: 39,

		CmdPassedFg: 18,
		CmdPassedBg: 7,
		CmdFailedFg: 254,
		CmdFailedBg: 124,

		SvnChangesFg: 148,
		SvnChangesBg: 22, // dark green

		GitAheadFg:      240,
		GitAheadBg:      250,
		GitBehindFg:     240,
		GitBehindBg:     251,
		GitStagedFg:     22,
		GitStagedBg:     15,
		GitNotStagedFg:  130,
		GitNotStagedBg:  15,
		GitUntrackedFg:  52,
		GitUntrackedBg:  15,
		GitConflictedFg: 9,
		GitConflictedBg: 15,
		GitStashedFg:    15,
		GitStashedBg:    20,

		VirtualEnvFg: 35, // a mid-tone green
		VirtualEnvBg: 254,

		VirtualGoFg: 35,
		VirtualGoBg: 254,

		PerlbrewFg: 20, // a mid-tone blue
		PerlbrewBg: 15,

		TimeFg: 236,
		TimeBg: 15,

		LoadFg:           15,
		LoadBg:           22,
		LoadHighBg:       161,
		LoadAvgValue:     5,
		LoadThresholdBad: 1.0,

		NixShellFg: 69, // a light blue
		NixShellBg: 254,

		HostnameColorizedFgMap: map[uint8]uint8{
			0:   250,
			1:   250,
			2:   120,
			3:   228,
			4:   250,
			5:   250,
			6:   123,
			7:   238,
			8:   0,
			9:   0,
			10:  0,
			11:  0,
			12:  250,
			13:  0,
			14:  0,
			15:  242,
			16:  250,
			17:  250,
			18:  250,
			19:  189,
			20:  254,
			21:  250,
			22:  83,
			23:  87,
			24:  117,
			25:  188,
			26:  254,
			27:  0,
			28:  120,
			29:  122,
			30:  123,
			31:  159,
			32:  255,
			33:  0,
			34:  157,
			35:  158,
			36:  159,
			37:  159,
			38:  195,
			39:  0,
			40:  194,
			41:  194,
			42:  195,
			43:  195,
			44:  195,
			45:  0,
			46:  0,
			47:  0,
			48:  0,
			49:  0,
			50:  0,
			51:  0,
			52:  250,
			53:  250,
			54:  250,
			55:  189,
			56:  254,
			57:  250,
			58:  227,
			59:  253,
			60:  255,
			61:  0,
			62:  233,
			63:  17,
			64:  192,
			65:  255,
			66:  195,
			67:  232,
			68:  233,
			69:  17,
			70:  193,
			71:  232,
			72:  232,
			73:  232,
			74:  234,
			75:  236,
			76:  194,
			77:  235,
			78:  235,
			79:  235,
			80:  235,
			81:  237,
			82:  0,
			83:  237,
			84:  237,
			85:  237,
			86:  237,
			87:  237,
			88:  250,
			89:  250,
			90:  250,
			91:  189,
			92:  254,
			93:  0,
			94:  222,
			95:  255,
			96:  255,
			97:  232,
			98:  233,
			99:  17,
			100: 228,
			101: 15,
			102: 232,
			103: 233,
			104: 17,
			105: 18,
			106: 229,
			107: 232,
			108: 234,
			109: 234,
			110: 236,
			111: 54,
			112: 230,
			113: 235,
			114: 22,
			115: 237,
			116: 238,
			117: 238,
			118: 0,
			119: 237,
			120: 22,
			121: 23,
			122: 23,
			123: 23,
			124: 252,
			125: 252,
			126: 189,
			127: 189,
			128: 254,
			129: 0,
			130: 223,
			131: 232,
			132: 232,
			133: 232,
			134: 233,
			135: 17,
			136: 229,
			137: 232,
			138: 233,
			139: 234,
			140: 53,
			141: 18,
			142: 229,
			143: 232,
			144: 234,
			145: 236,
			146: 17,
			147: 19,
			148: 230,
			149: 235,
			150: 238,
			151: 22,
			152: 23,
			153: 24,
			154: 0,
			155: 237,
			156: 22,
			157: 2,
			158: 29,
			159: 6,
			160: 254,
			161: 254,
			162: 254,
			163: 254,
			164: 254,
			165: 0,
			166: 255,
			167: 233,
			168: 233,
			169: 234,
			170: 234,
			171: 235,
			172: 230,
			173: 234,
			174: 52,
			175: 235,
			176: 53,
			177: 53,
			178: 230,
			179: 235,
			180: 236,
			181: 52,
			182: 53,
			183: 55,
			184: 230,
			185: 235,
			186: 238,
			187: 58,
			188: 240,
			189: 20,
			190: 0,
			191: 238,
			192: 58,
			193: 64,
			194: 35,
			195: 66,
			196: 0,
			197: 0,
			198: 0,
			199: 0,
			200: 0,
			201: 0,
			202: 0,
			203: 235,
			204: 235,
			205: 235,
			206: 235,
			207: 53,
			208: 0,
			209: 236,
			210: 52,
			211: 237,
			212: 53,
			213: 53,
			214: 0,
			215: 236,
			216: 238,
			217: 1,
			218: 89,
			219: 5,
			220: 0,
			221: 237,
			222: 58,
			223: 95,
			224: 131,
			225: 126,
			226: 0,
			227: 238,
			228: 58,
			229: 3,
			230: 143,
			231: 242,
			232: 250,
			233: 250,
			234: 250,
			235: 250,
			236: 250,
			237: 250,
			238: 251,
			239: 252,
			240: 188,
			241: 254,
			242: 254,
			243: 255,
			244: 0,
			245: 232,
			246: 233,
			247: 234,
			248: 235,
			249: 236,
			250: 237,
			251: 238,
			252: 239,
			253: 240,
			254: 242,
			255: 243,
		},
	},
	"solarized-dark16": {
		Reset:              8,
		UsernameFg:         15,
		UsernameBg:         4,
		UsernameRootBg:     1,
		HostnameFg:         15,
		HostnameBg:         0,
		HomeSpecialDisplay: false,
		HomeFg:             15,
		HomeBg:             4,
		PathFg:             15,
		PathBg:             10,
		CwdFg:              15,
		SeparatorFg:        15,
		ReadonlyFg:         8,
		ReadonlyBg:         1,
		SshFg:              8,
		SshBg:              9,
		DockerMachineFg:    13,
		DockerMachineBg:    55,
		DotEnvFg:           15,
		DotEnvBg:           55,
		RepoCleanFg:        15,
		RepoCleanBg:        3,
		RepoDirtyFg:        15,
		RepoDirtyBg:        5,
		JobsFg:             4,
		JobsBg:             0,
		CmdPassedFg:        15,
		CmdPassedBg:        0,
		CmdFailedFg:        15,
		CmdFailedBg:        5,
		SvnChangesFg:       2,
		SvnChangesBg:       3,
		GitAheadFg:         14,
		GitAheadBg:         10,
		GitBehindFg:        14,
		GitBehindBg:        10,
		GitStagedFg:        15,
		GitStagedBg:        2,
		GitNotStagedFg:     15,
		GitNotStagedBg:     9,
		GitUntrackedFg:     15,
		GitUntrackedBg:     1,
		GitConflictedFg:    15,
		GitConflictedBg:    1,
		GitStashedFg:       15,
		GitStashedBg:       4,
		VirtualEnvFg:       8,
		VirtualEnvBg:       6,
		PerlbrewFg:         8,
		PerlbrewBg:         4,
		TimeFg:             15,
		TimeBg:             0,
		ShellVarFg:         1,
		ShellVarBg:         11,
		NodeFg:             15,
		NodeBg:             40,
		LoadFg:             15,
		LoadBg:             2,
		LoadHighBg:         5,
		LoadAvgValue:       5,
		LoadThresholdBad:   1.0,
		NixShellFg:         0,
		NixShellBg:         4,
		HostnameColorizedFgMap: map[uint8]uint8{
			0:   14,
			1:   14,
			2:   120,
			3:   228,
			4:   14,
			5:   14,
			6:   123,
			7:   0,
			8:   8,
			9:   8,
			10:  8,
			11:  8,
			12:  14,
			13:  8,
			14:  8,
			15:  242,
			16:  14,
			17:  14,
			18:  14,
			19:  189,
			20:  8,
			21:  14,
			22:  83,
			23:  87,
			24:  117,
			25:  188,
			26:  8,
			27:  8,
			28:  120,
			29:  122,
			30:  123,
			31:  159,
			32:  8,
			33:  8,
			34:  157,
			35:  158,
			36:  159,
			37:  159,
			38:  195,
			39:  8,
			40:  194,
			41:  194,
			42:  195,
			43:  195,
			44:  195,
			45:  8,
			46:  8,
			47:  8,
			48:  8,
			49:  8,
			50:  8,
			51:  8,
			52:  14,
			53:  14,
			54:  14,
			55:  189,
			56:  8,
			57:  14,
			58:  227,
			59:  253,
			60:  8,
			61:  8,
			62:  233,
			63:  17,
			64:  192,
			65:  8,
			66:  195,
			67:  232,
			68:  233,
			69:  17,
			70:  193,
			71:  232,
			72:  232,
			73:  232,
			74:  8,
			75:  0,
			76:  194,
			77:  0,
			78:  0,
			79:  0,
			80:  0,
			81:  0,
			82:  8,
			83:  0,
			84:  0,
			85:  0,
			86:  0,
			87:  0,
			88:  14,
			89:  14,
			90:  14,
			91:  189,
			92:  8,
			93:  8,
			94:  222,
			95:  8,
			96:  8,
			97:  232,
			98:  233,
			99:  17,
			100: 228,
			101: 15,
			102: 232,
			103: 233,
			104: 17,
			105: 18,
			106: 229,
			107: 232,
			108: 8,
			109: 8,
			110: 0,
			111: 54,
			112: 15,
			113: 0,
			114: 2,
			115: 0,
			116: 0,
			117: 0,
			118: 8,
			119: 0,
			120: 2,
			121: 23,
			122: 23,
			123: 23,
			124: 252,
			125: 252,
			126: 189,
			127: 189,
			128: 8,
			129: 8,
			130: 223,
			131: 232,
			132: 232,
			133: 232,
			134: 233,
			135: 17,
			136: 229,
			137: 232,
			138: 233,
			139: 8,
			140: 53,
			141: 18,
			142: 229,
			143: 232,
			144: 8,
			145: 0,
			146: 17,
			147: 19,
			148: 15,
			149: 0,
			150: 0,
			151: 2,
			152: 23,
			153: 24,
			154: 8,
			155: 0,
			156: 2,
			157: 2,
			158: 29,
			159: 6,
			160: 8,
			161: 8,
			162: 8,
			163: 8,
			164: 8,
			165: 8,
			166: 8,
			167: 233,
			168: 233,
			169: 8,
			170: 8,
			171: 0,
			172: 15,
			173: 8,
			174: 1,
			175: 0,
			176: 53,
			177: 53,
			178: 15,
			179: 0,
			180: 0,
			181: 1,
			182: 53,
			183: 55,
			184: 15,
			185: 0,
			186: 0,
			187: 58,
			188: 10,
			189: 4,
			190: 8,
			191: 0,
			192: 58,
			193: 2,
			194: 35,
			195: 66,
			196: 8,
			197: 8,
			198: 8,
			199: 8,
			200: 8,
			201: 8,
			202: 8,
			203: 0,
			204: 0,
			205: 0,
			206: 0,
			207: 53,
			208: 8,
			209: 0,
			210: 1,
			211: 0,
			212: 53,
			213: 53,
			214: 8,
			215: 0,
			216: 0,
			217: 1,
			218: 89,
			219: 5,
			220: 8,
			221: 0,
			222: 58,
			223: 95,
			224: 131,
			225: 126,
			226: 8,
			227: 0,
			228: 58,
			229: 3,
			230: 143,
			231: 242,
			232: 14,
			233: 14,
			234: 14,
			235: 14,
			236: 14,
			237: 14,
			238: 251,
			239: 252,
			240: 188,
			241: 8,
			242: 8,
			243: 8,
			244: 8,
			245: 232,
			246: 233,
			247: 8,
			248: 0,
			249: 0,
			250: 0,
			251: 0,
			252: 239,
			253: 10,
			254: 242,
			255: 243,
		},
	},
	"solarized-light16": {
		Reset:              0,
		UsernameFg:         15,
		UsernameBg:         4,
		UsernameRootBg:     1,
		HostnameFg:         15,
		HostnameBg:         7,
		HomeSpecialDisplay: false,
		HomeFg:             15,
		HomeBg:             4,
		PathFg:             15,
		PathBg:             10,
		CwdFg:              15,
		SeparatorFg:        15,
		ReadonlyFg:         8,
		ReadonlyBg:         1,
		SshFg:              8,
		SshBg:              9,
		DockerMachineFg:    13,
		DockerMachineBg:    55,
		DotEnvFg:           15,
		DotEnvBg:           55,
		RepoCleanFg:        15,
		RepoCleanBg:        3,
		RepoDirtyFg:        15,
		RepoDirtyBg:        5,
		JobsFg:             4,
		JobsBg:             0,
		CmdPassedFg:        10,
		CmdPassedBg:        7,
		CmdFailedFg:        15,
		CmdFailedBg:        5,
		SvnChangesFg:       2,
		SvnChangesBg:       3,
		GitAheadFg:         14,
		GitAheadBg:         10,
		GitBehindFg:        14,
		GitBehindBg:        10,
		GitStagedFg:        15,
		GitStagedBg:        2,
		GitNotStagedFg:     15,
		GitNotStagedBg:     9,
		GitUntrackedFg:     15,
		GitUntrackedBg:     1,
		GitConflictedFg:    15,
		GitConflictedBg:    1,
		GitStashedFg:       15,
		GitStashedBg:       4,
		VirtualEnvFg:       8,
		VirtualEnvBg:       6,
		PerlbrewFg:         8,
		PerlbrewBg:         4,
		TimeFg:             15,
		TimeBg:             0,
		ShellVarFg:         1,
		ShellVarBg:         11,
		NodeFg:             15,
		NodeBg:             40,
		LoadFg:             15,
		LoadBg:             2,
		LoadHighBg:         5,
		LoadAvgValue:       5,
		LoadThresholdBad:   1.0,
		NixShellFg:         0,
		NixShellBg:         7,
		HostnameColorizedFgMap: map[uint8]uint8{
			0:   14,
			1:   14,
			2:   120,
			3:   228,
			4:   14,
			5:   14,
			6:   123,
			7:   0,
			8:   8,
			9:   8,
			10:  8,
			11:  8,
			12:  14,
			13:  8,
			14:  8,
			15:  242,
			16:  14,
			17:  14,
			18:  14,
			19:  189,
			20:  8,
			21:  14,
			22:  83,
			23:  87,
			24:  117,
			25:  188,
			26:  8,
			27:  8,
			28:  120,
			29:  122,
			30:  123,
			31:  159,
			32:  8,
			33:  8,
			34:  157,
			35:  158,
			36:  159,
			37:  159,
			38:  195,
			39:  8,
			40:  194,
			41:  194,
			42:  195,
			43:  195,
			44:  195,
			45:  8,
			46:  8,
			47:  8,
			48:  8,
			49:  8,
			50:  8,
			51:  8,
			52:  14,
			53:  14,
			54:  14,
			55:  189,
			56:  8,
			57:  14,
			58:  227,
			59:  253,
			60:  8,
			61:  8,
			62:  233,
			63:  17,
			64:  192,
			65:  8,
			66:  195,
			67:  232,
			68:  233,
			69:  17,
			70:  193,
			71:  232,
			72:  232,
			73:  232,
			74:  8,
			75:  0,
			76:  194,
			77:  0,
			78:  0,
			79:  0,
			80:  0,
			81:  0,
			82:  8,
			83:  0,
			84:  0,
			85:  0,
			86:  0,
			87:  0,
			88:  14,
			89:  14,
			90:  14,
			91:  189,
			92:  8,
			93:  8,
			94:  222,
			95:  8,
			96:  8,
			97:  232,
			98:  233,
			99:  17,
			100: 228,
			101: 15,
			102: 232,
			103: 233,
			104: 17,
			105: 18,
			106: 229,
			107: 232,
			108: 8,
			109: 8,
			110: 0,
			111: 54,
			112: 15,
			113: 0,
			114: 2,
			115: 0,
			116: 0,
			117: 0,
			118: 8,
			119: 0,
			120: 2,
			121: 23,
			122: 23,
			123: 23,
			124: 252,
			125: 252,
			126: 189,
			127: 189,
			128: 8,
			129: 8,
			130: 223,
			131: 232,
			132: 232,
			133: 232,
			134: 233,
			135: 17,
			136: 229,
			137: 232,
			138: 233,
			139: 8,
			140: 53,
			141: 18,
			142: 229,
			143: 232,
			144: 8,
			145: 0,
			146: 17,
			147: 19,
			148: 15,
			149: 0,
			150: 0,
			151: 2,
			152: 23,
			153: 24,
			154: 8,
			155: 0,
			156: 2,
			157: 2,
			158: 29,
			159: 6,
			160: 8,
			161: 8,
			162: 8,
			163: 8,
			164: 8,
			165: 8,
			166: 8,
			167: 233,
			168: 233,
			169: 8,
			170: 8,
			171: 0,
			172: 15,
			173: 8,
			174: 1,
			175: 0,
			176: 53,
			177: 53,
			178: 15,
			179: 0,
			180: 0,
			181: 1,
			182: 53,
			183: 55,
			184: 15,
			185: 0,
			186: 0,
			187: 58,
			188: 10,
			189: 4,
			190: 8,
			191: 0,
			192: 58,
			193: 2,
			194: 35,
			195: 66,
			196: 8,
			197: 8,
			198: 8,
			199: 8,
			200: 8,
			201: 8,
			202: 8,
			203: 0,
			204: 0,
			205: 0,
			206: 0,
			207: 53,
			208: 8,
			209: 0,
			210: 1,
			211: 0,
			212: 53,
			213: 53,
			214: 8,
			215: 0,
			216: 0,
			217: 1,
			218: 89,
			219: 5,
			220: 8,
			221: 0,
			222: 58,
			223: 95,
			224: 131,
			225: 126,
			226: 8,
			227: 0,
			228: 58,
			229: 3,
			230: 143,
			231: 242,
			232: 14,
			233: 14,
			234: 14,
			235: 14,
			236: 14,
			237: 14,
			238: 251,
			239: 252,
			240: 188,
			241: 8,
			242: 8,
			243: 8,
			244: 8,
			245: 232,
			246: 233,
			247: 8,
			248: 0,
			249: 0,
			250: 0,
			251: 0,
			252: 239,
			253: 10,
			254: 242,
			255: 243,
		},
	},
}
