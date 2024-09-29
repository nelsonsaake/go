package xls2

import "github.com/xuri/excelize/v2"

var (
	FillSlate50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f8fafc",
			"#f8fafc",
		},
	}

	FillSlate100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f1f5f9",
			"#f1f5f9",
		},
	}

	FillSlate200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#e2e8f0",
			"#e2e8f0",
		},
	}

	FillSlate300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#cbd5e1",
			"#cbd5e1",
		},
	}

	FillSlate400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#94a3b8",
			"#94a3b8",
		},
	}

	FillSlate500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#64748b",
			"#64748b",
		},
	}

	FillSlate600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#475569",
			"#475569",
		},
	}

	FillSlate700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#334155",
			"#334155",
		},
	}

	FillSlate800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#1e293b",
			"#1e293b",
		},
	}

	FillSlate900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#0f172a",
			"#0f172a",
		},
	}

	FillSlate950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#020617",
			"#020617",
		},
	}

	FillGray50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f9fafb",
			"#f9fafb",
		},
	}

	FillGray100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f3f4f6",
			"#f3f4f6",
		},
	}

	FillGray200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#e5e7eb",
			"#e5e7eb",
		},
	}

	FillGray300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#d1d5db",
			"#d1d5db",
		},
	}

	FillGray400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#9ca3af",
			"#9ca3af",
		},
	}

	FillGray500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#6b7280",
			"#6b7280",
		},
	}

	FillGray600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#4b5563",
			"#4b5563",
		},
	}

	FillGray700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#374151",
			"#374151",
		},
	}

	FillGray800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#1f2937",
			"#1f2937",
		},
	}

	FillGray900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#111827",
			"#111827",
		},
	}

	FillGray950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#030712",
			"#030712",
		},
	}

	FillZinc50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fafafa",
			"#fafafa",
		},
	}

	FillZinc100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f4f4f5",
			"#f4f4f5",
		},
	}

	FillZinc200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#e4e4e7",
			"#e4e4e7",
		},
	}

	FillZinc300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#d4d4d8",
			"#d4d4d8",
		},
	}

	FillZinc400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#a1a1aa",
			"#a1a1aa",
		},
	}

	FillZinc500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#71717a",
			"#71717a",
		},
	}

	FillZinc600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#52525b",
			"#52525b",
		},
	}

	FillZinc700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#3f3f46",
			"#3f3f46",
		},
	}

	FillZinc800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#27272a",
			"#27272a",
		},
	}

	FillZinc900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#18181b",
			"#18181b",
		},
	}

	FillZinc950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#09090b",
			"#09090b",
		},
	}

	FillNeutral50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fafafa",
			"#fafafa",
		},
	}

	FillNeutral100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f5f5f5",
			"#f5f5f5",
		},
	}

	FillNeutral200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#e5e5e5",
			"#e5e5e5",
		},
	}

	FillNeutral300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#d4d4d4",
			"#d4d4d4",
		},
	}

	FillNeutral400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#a3a3a3",
			"#a3a3a3",
		},
	}

	FillNeutral500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#737373",
			"#737373",
		},
	}

	FillNeutral600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#525252",
			"#525252",
		},
	}

	FillNeutral700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#404040",
			"#404040",
		},
	}

	FillNeutral800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#262626",
			"#262626",
		},
	}

	FillNeutral900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#171717",
			"#171717",
		},
	}

	FillNeutral950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#0a0a0a",
			"#0a0a0a",
		},
	}

	FillStone50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fafaf9",
			"#fafaf9",
		},
	}

	FillStone100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f5f5f4",
			"#f5f5f4",
		},
	}

	FillStone200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#e7e5e4",
			"#e7e5e4",
		},
	}

	FillStone300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#d6d3d1",
			"#d6d3d1",
		},
	}

	FillStone400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#a8a29e",
			"#a8a29e",
		},
	}

	FillStone500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#78716c",
			"#78716c",
		},
	}

	FillStone600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#57534e",
			"#57534e",
		},
	}

	FillStone700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#44403c",
			"#44403c",
		},
	}

	FillStone800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#292524",
			"#292524",
		},
	}

	FillStone900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#1c1917",
			"#1c1917",
		},
	}

	FillStone950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#0c0a09",
			"#0c0a09",
		},
	}

	FillRed50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fef2f2",
			"#fef2f2",
		},
	}

	FillRed100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fee2e2",
			"#fee2e2",
		},
	}

	FillRed200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fecaca",
			"#fecaca",
		},
	}

	FillRed300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fca5a5",
			"#fca5a5",
		},
	}

	FillRed400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f87171",
			"#f87171",
		},
	}

	FillRed500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ef4444",
			"#ef4444",
		},
	}

	FillRed600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#dc2626",
			"#dc2626",
		},
	}

	FillRed700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#b91c1c",
			"#b91c1c",
		},
	}

	FillRed800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#991b1b",
			"#991b1b",
		},
	}

	FillRed900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#7f1d1d",
			"#7f1d1d",
		},
	}

	FillRed950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#450a0a",
			"#450a0a",
		},
	}

	FillOrange50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fff7ed",
			"#fff7ed",
		},
	}

	FillOrange100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ffedd5",
			"#ffedd5",
		},
	}

	FillOrange200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fed7aa",
			"#fed7aa",
		},
	}

	FillOrange300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fdba74",
			"#fdba74",
		},
	}

	FillOrange400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fb923c",
			"#fb923c",
		},
	}

	FillOrange500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f97316",
			"#f97316",
		},
	}

	FillOrange600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ea580c",
			"#ea580c",
		},
	}

	FillOrange700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#c2410c",
			"#c2410c",
		},
	}

	FillOrange800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#9a3412",
			"#9a3412",
		},
	}

	FillOrange900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#7c2d12",
			"#7c2d12",
		},
	}

	FillOrange950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#431407",
			"#431407",
		},
	}

	FillAmber50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fffbeb",
			"#fffbeb",
		},
	}

	FillAmber100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fef3c7",
			"#fef3c7",
		},
	}

	FillAmber200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fde68a",
			"#fde68a",
		},
	}

	FillAmber300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fcd34d",
			"#fcd34d",
		},
	}

	FillAmber400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fbbf24",
			"#fbbf24",
		},
	}

	FillAmber500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f59e0b",
			"#f59e0b",
		},
	}

	FillAmber600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#d97706",
			"#d97706",
		},
	}

	FillAmber700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#b45309",
			"#b45309",
		},
	}

	FillAmber800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#92400e",
			"#92400e",
		},
	}

	FillAmber900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#78350f",
			"#78350f",
		},
	}

	FillAmber950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#451a03",
			"#451a03",
		},
	}

	FillYellow50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fefce8",
			"#fefce8",
		},
	}

	FillYellow100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fef9c3",
			"#fef9c3",
		},
	}

	FillYellow200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fef08a",
			"#fef08a",
		},
	}

	FillYellow300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fde047",
			"#fde047",
		},
	}

	FillYellow400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#facc15",
			"#facc15",
		},
	}

	FillYellow500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#eab308",
			"#eab308",
		},
	}

	FillYellow600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ca8a04",
			"#ca8a04",
		},
	}

	FillYellow700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#a16207",
			"#a16207",
		},
	}

	FillYellow800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#854d0e",
			"#854d0e",
		},
	}

	FillYellow900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#713f12",
			"#713f12",
		},
	}

	FillYellow950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#422006",
			"#422006",
		},
	}

	FillLime50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f7fee7",
			"#f7fee7",
		},
	}

	FillLime100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ecfccb",
			"#ecfccb",
		},
	}

	FillLime200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#d9f99d",
			"#d9f99d",
		},
	}

	FillLime300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#bef264",
			"#bef264",
		},
	}

	FillLime400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#a3e635",
			"#a3e635",
		},
	}

	FillLime500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#84cc16",
			"#84cc16",
		},
	}

	FillLime600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#65a30d",
			"#65a30d",
		},
	}

	FillLime700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#4d7c0f",
			"#4d7c0f",
		},
	}

	FillLime800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#3f6212",
			"#3f6212",
		},
	}

	FillLime900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#365314",
			"#365314",
		},
	}

	FillLime950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#1a2e05",
			"#1a2e05",
		},
	}

	FillGreen50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f0fdf4",
			"#f0fdf4",
		},
	}

	FillGreen100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#dcfce7",
			"#dcfce7",
		},
	}

	FillGreen200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#bbf7d0",
			"#bbf7d0",
		},
	}

	FillGreen300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#86efac",
			"#86efac",
		},
	}

	FillGreen400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#4ade80",
			"#4ade80",
		},
	}

	FillGreen500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#22c55e",
			"#22c55e",
		},
	}

	FillGreen600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#16a34a",
			"#16a34a",
		},
	}

	FillGreen700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#15803d",
			"#15803d",
		},
	}

	FillGreen800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#166534",
			"#166534",
		},
	}

	FillGreen900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#14532d",
			"#14532d",
		},
	}

	FillGreen950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#052e16",
			"#052e16",
		},
	}

	FillEmerald50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ecfdf5",
			"#ecfdf5",
		},
	}

	FillEmerald100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#d1fae5",
			"#d1fae5",
		},
	}

	FillEmerald200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#a7f3d0",
			"#a7f3d0",
		},
	}

	FillEmerald300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#6ee7b7",
			"#6ee7b7",
		},
	}

	FillEmerald400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#34d399",
			"#34d399",
		},
	}

	FillEmerald500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#10b981",
			"#10b981",
		},
	}

	FillEmerald600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#059669",
			"#059669",
		},
	}

	FillEmerald700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#047857",
			"#047857",
		},
	}

	FillEmerald800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#065f46",
			"#065f46",
		},
	}

	FillEmerald900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#064e3b",
			"#064e3b",
		},
	}

	FillEmerald950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#022c22",
			"#022c22",
		},
	}

	FillTeal50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f0fdfa",
			"#f0fdfa",
		},
	}

	FillTeal100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ccfbf1",
			"#ccfbf1",
		},
	}

	FillTeal200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#99f6e4",
			"#99f6e4",
		},
	}

	FillTeal300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#5eead4",
			"#5eead4",
		},
	}

	FillTeal400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#2dd4bf",
			"#2dd4bf",
		},
	}

	FillTeal500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#14b8a6",
			"#14b8a6",
		},
	}

	FillTeal600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#0d9488",
			"#0d9488",
		},
	}

	FillTeal700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#0f766e",
			"#0f766e",
		},
	}

	FillTeal800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#115e59",
			"#115e59",
		},
	}

	FillTeal900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#134e4a",
			"#134e4a",
		},
	}

	FillTeal950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#042f2e",
			"#042f2e",
		},
	}

	FillCyan50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ecfeff",
			"#ecfeff",
		},
	}

	FillCyan100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#cffafe",
			"#cffafe",
		},
	}

	FillCyan200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#a5f3fc",
			"#a5f3fc",
		},
	}

	FillCyan300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#67e8f9",
			"#67e8f9",
		},
	}

	FillCyan400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#22d3ee",
			"#22d3ee",
		},
	}

	FillCyan500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#06b6d4",
			"#06b6d4",
		},
	}

	FillCyan600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#0891b2",
			"#0891b2",
		},
	}

	FillCyan700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#0e7490",
			"#0e7490",
		},
	}

	FillCyan800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#155e75",
			"#155e75",
		},
	}

	FillCyan900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#164e63",
			"#164e63",
		},
	}

	FillCyan950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#083344",
			"#083344",
		},
	}

	FillSky50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f0f9ff",
			"#f0f9ff",
		},
	}

	FillSky100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#e0f2fe",
			"#e0f2fe",
		},
	}

	FillSky200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#bae6fd",
			"#bae6fd",
		},
	}

	FillSky300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#7dd3fc",
			"#7dd3fc",
		},
	}

	FillSky400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#38bdf8",
			"#38bdf8",
		},
	}

	FillSky500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#0ea5e9",
			"#0ea5e9",
		},
	}

	FillSky600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#0284c7",
			"#0284c7",
		},
	}

	FillSky700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#0369a1",
			"#0369a1",
		},
	}

	FillSky800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#075985",
			"#075985",
		},
	}

	FillSky900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#0c4a6e",
			"#0c4a6e",
		},
	}

	FillSky950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#082f49",
			"#082f49",
		},
	}

	FillBlue50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#eff6ff",
			"#eff6ff",
		},
	}

	FillBlue100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#dbeafe",
			"#dbeafe",
		},
	}

	FillBlue200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#bfdbfe",
			"#bfdbfe",
		},
	}

	FillBlue300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#93c5fd",
			"#93c5fd",
		},
	}

	FillBlue400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#60a5fa",
			"#60a5fa",
		},
	}

	FillBlue500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#3b82f6",
			"#3b82f6",
		},
	}

	FillBlue600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#2563eb",
			"#2563eb",
		},
	}

	FillBlue700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#1d4ed8",
			"#1d4ed8",
		},
	}

	FillBlue800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#1e40af",
			"#1e40af",
		},
	}

	FillBlue900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#1e3a8a",
			"#1e3a8a",
		},
	}

	FillBlue950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#172554",
			"#172554",
		},
	}

	FillIndigo50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#eef2ff",
			"#eef2ff",
		},
	}

	FillIndigo100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#e0e7ff",
			"#e0e7ff",
		},
	}

	FillIndigo200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#c7d2fe",
			"#c7d2fe",
		},
	}

	FillIndigo300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#a5b4fc",
			"#a5b4fc",
		},
	}

	FillIndigo400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#818cf8",
			"#818cf8",
		},
	}

	FillIndigo500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#6366f1",
			"#6366f1",
		},
	}

	FillIndigo600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#4f46e5",
			"#4f46e5",
		},
	}

	FillIndigo700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#4338ca",
			"#4338ca",
		},
	}

	FillIndigo800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#3730a3",
			"#3730a3",
		},
	}

	FillIndigo900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#312e81",
			"#312e81",
		},
	}

	FillIndigo950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#1e1b4b",
			"#1e1b4b",
		},
	}

	FillViolet50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f5f3ff",
			"#f5f3ff",
		},
	}

	FillViolet100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ede9fe",
			"#ede9fe",
		},
	}

	FillViolet200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ddd6fe",
			"#ddd6fe",
		},
	}

	FillViolet300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#c4b5fd",
			"#c4b5fd",
		},
	}

	FillViolet400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#a78bfa",
			"#a78bfa",
		},
	}

	FillViolet500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#8b5cf6",
			"#8b5cf6",
		},
	}

	FillViolet600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#7c3aed",
			"#7c3aed",
		},
	}

	FillViolet700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#6d28d9",
			"#6d28d9",
		},
	}

	FillViolet800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#5b21b6",
			"#5b21b6",
		},
	}

	FillViolet900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#4c1d95",
			"#4c1d95",
		},
	}

	FillViolet950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#2e1065",
			"#2e1065",
		},
	}

	FillPurple50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#faf5ff",
			"#faf5ff",
		},
	}

	FillPurple100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f3e8ff",
			"#f3e8ff",
		},
	}

	FillPurple200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#e9d5ff",
			"#e9d5ff",
		},
	}

	FillPurple300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#d8b4fe",
			"#d8b4fe",
		},
	}

	FillPurple400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#c084fc",
			"#c084fc",
		},
	}

	FillPurple500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#a855f7",
			"#a855f7",
		},
	}

	FillPurple600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#9333ea",
			"#9333ea",
		},
	}

	FillPurple700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#7e22ce",
			"#7e22ce",
		},
	}

	FillPurple800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#6b21a8",
			"#6b21a8",
		},
	}

	FillPurple900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#581c87",
			"#581c87",
		},
	}

	FillPurple950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#3b0764",
			"#3b0764",
		},
	}

	FillFuchsia50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fdf4ff",
			"#fdf4ff",
		},
	}

	FillFuchsia100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fae8ff",
			"#fae8ff",
		},
	}

	FillFuchsia200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f5d0fe",
			"#f5d0fe",
		},
	}

	FillFuchsia300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f0abfc",
			"#f0abfc",
		},
	}

	FillFuchsia400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#e879f9",
			"#e879f9",
		},
	}

	FillFuchsia500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#d946ef",
			"#d946ef",
		},
	}

	FillFuchsia600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#c026d3",
			"#c026d3",
		},
	}

	FillFuchsia700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#a21caf",
			"#a21caf",
		},
	}

	FillFuchsia800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#86198f",
			"#86198f",
		},
	}

	FillFuchsia900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#701a75",
			"#701a75",
		},
	}

	FillFuchsia950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#4a044e",
			"#4a044e",
		},
	}

	FillPink50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fdf2f8",
			"#fdf2f8",
		},
	}

	FillPink100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fce7f3",
			"#fce7f3",
		},
	}

	FillPink200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fbcfe8",
			"#fbcfe8",
		},
	}

	FillPink300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f9a8d4",
			"#f9a8d4",
		},
	}

	FillPink400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f472b6",
			"#f472b6",
		},
	}

	FillPink500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ec4899",
			"#ec4899",
		},
	}

	FillPink600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#db2777",
			"#db2777",
		},
	}

	FillPink700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#be185d",
			"#be185d",
		},
	}

	FillPink800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#9d174d",
			"#9d174d",
		},
	}

	FillPink900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#831843",
			"#831843",
		},
	}

	FillPink950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#500724",
			"#500724",
		},
	}

	FillRose50 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fff1f2",
			"#fff1f2",
		},
	}

	FillRose100 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#ffe4e6",
			"#ffe4e6",
		},
	}

	FillRose200 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fecdd3",
			"#fecdd3",
		},
	}

	FillRose300 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fda4af",
			"#fda4af",
		},
	}

	FillRose400 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#fb7185",
			"#fb7185",
		},
	}

	FillRose500 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#f43f5e",
			"#f43f5e",
		},
	}

	FillRose600 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#e11d48",
			"#e11d48",
		},
	}

	FillRose700 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#be123c",
			"#be123c",
		},
	}

	FillRose800 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#9f1239",
			"#9f1239",
		},
	}

	FillRose900 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#881337",
			"#881337",
		},
	}

	FillRose950 = excelize.Fill{
		Type: "gradient",
		Color: []string{
			"#4c0519",
			"#4c0519",
		},
	}
)
