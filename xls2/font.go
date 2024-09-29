package xls2

import "github.com/xuri/excelize/v2"

var (
	FontSlate50 = &excelize.Font{
		Color: "#f8fafc",
	}

	FontSlate50Italic = &excelize.Font{
		Color:  "#f8fafc",
		Italic: true,
	}

	FontSlate100 = &excelize.Font{
		Color: "#f1f5f9",
	}

	FontSlate100Italic = &excelize.Font{
		Color:  "#f1f5f9",
		Italic: true,
	}

	FontSlate200 = &excelize.Font{
		Color: "#e2e8f0",
	}

	FontSlate200Italic = &excelize.Font{
		Color:  "#e2e8f0",
		Italic: true,
	}

	FontSlate300 = &excelize.Font{
		Color: "#cbd5e1",
	}

	FontSlate300Italic = &excelize.Font{
		Color:  "#cbd5e1",
		Italic: true,
	}

	FontSlate400 = &excelize.Font{
		Color: "#94a3b8",
	}

	FontSlate400Italic = &excelize.Font{
		Color:  "#94a3b8",
		Italic: true,
	}

	FontSlate500 = &excelize.Font{
		Color: "#64748b",
	}

	FontSlate500Italic = &excelize.Font{
		Color:  "#64748b",
		Italic: true,
	}

	FontSlate600 = &excelize.Font{
		Color: "#475569",
	}

	FontSlate600Italic = &excelize.Font{
		Color:  "#475569",
		Italic: true,
	}

	FontSlate700 = &excelize.Font{
		Color: "#334155",
	}

	FontSlate700Italic = &excelize.Font{
		Color:  "#334155",
		Italic: true,
	}

	FontSlate800 = &excelize.Font{
		Color: "#1e293b",
	}

	FontSlate800Italic = &excelize.Font{
		Color:  "#1e293b",
		Italic: true,
	}

	FontSlate900 = &excelize.Font{
		Color: "#0f172a",
	}

	FontSlate900Italic = &excelize.Font{
		Color:  "#0f172a",
		Italic: true,
	}

	FontSlate950 = &excelize.Font{
		Color: "#020617",
	}

	FontSlate950Italic = &excelize.Font{
		Color:  "#020617",
		Italic: true,
	}

	FontGray50 = &excelize.Font{
		Color: "#f9fafb",
	}

	FontGray50Italic = &excelize.Font{
		Color:  "#f9fafb",
		Italic: true,
	}

	FontGray100 = &excelize.Font{
		Color: "#f3f4f6",
	}

	FontGray100Italic = &excelize.Font{
		Color:  "#f3f4f6",
		Italic: true,
	}

	FontGray200 = &excelize.Font{
		Color: "#e5e7eb",
	}

	FontGray200Italic = &excelize.Font{
		Color:  "#e5e7eb",
		Italic: true,
	}

	FontGray300 = &excelize.Font{
		Color: "#d1d5db",
	}

	FontGray300Italic = &excelize.Font{
		Color:  "#d1d5db",
		Italic: true,
	}

	FontGray400 = &excelize.Font{
		Color: "#9ca3af",
	}

	FontGray400Italic = &excelize.Font{
		Color:  "#9ca3af",
		Italic: true,
	}

	FontGray500 = &excelize.Font{
		Color: "#6b7280",
	}

	FontGray500Italic = &excelize.Font{
		Color:  "#6b7280",
		Italic: true,
	}

	FontGray600 = &excelize.Font{
		Color: "#4b5563",
	}

	FontGray600Italic = &excelize.Font{
		Color:  "#4b5563",
		Italic: true,
	}

	FontGray700 = &excelize.Font{
		Color: "#374151",
	}

	FontGray700Italic = &excelize.Font{
		Color:  "#374151",
		Italic: true,
	}

	FontGray800 = &excelize.Font{
		Color: "#1f2937",
	}

	FontGray800Italic = &excelize.Font{
		Color:  "#1f2937",
		Italic: true,
	}

	FontGray900 = &excelize.Font{
		Color: "#111827",
	}

	FontGray900Italic = &excelize.Font{
		Color:  "#111827",
		Italic: true,
	}

	FontGray950 = &excelize.Font{
		Color: "#030712",
	}

	FontGray950Italic = &excelize.Font{
		Color:  "#030712",
		Italic: true,
	}

	FontZinc50 = &excelize.Font{
		Color: "#fafafa",
	}

	FontZinc50Italic = &excelize.Font{
		Color:  "#fafafa",
		Italic: true,
	}

	FontZinc100 = &excelize.Font{
		Color: "#f4f4f5",
	}

	FontZinc100Italic = &excelize.Font{
		Color:  "#f4f4f5",
		Italic: true,
	}

	FontZinc200 = &excelize.Font{
		Color: "#e4e4e7",
	}

	FontZinc200Italic = &excelize.Font{
		Color:  "#e4e4e7",
		Italic: true,
	}

	FontZinc300 = &excelize.Font{
		Color: "#d4d4d8",
	}

	FontZinc300Italic = &excelize.Font{
		Color:  "#d4d4d8",
		Italic: true,
	}

	FontZinc400 = &excelize.Font{
		Color: "#a1a1aa",
	}

	FontZinc400Italic = &excelize.Font{
		Color:  "#a1a1aa",
		Italic: true,
	}

	FontZinc500 = &excelize.Font{
		Color: "#71717a",
	}

	FontZinc500Italic = &excelize.Font{
		Color:  "#71717a",
		Italic: true,
	}

	FontZinc600 = &excelize.Font{
		Color: "#52525b",
	}

	FontZinc600Italic = &excelize.Font{
		Color:  "#52525b",
		Italic: true,
	}

	FontZinc700 = &excelize.Font{
		Color: "#3f3f46",
	}

	FontZinc700Italic = &excelize.Font{
		Color:  "#3f3f46",
		Italic: true,
	}

	FontZinc800 = &excelize.Font{
		Color: "#27272a",
	}

	FontZinc800Italic = &excelize.Font{
		Color:  "#27272a",
		Italic: true,
	}

	FontZinc900 = &excelize.Font{
		Color: "#18181b",
	}

	FontZinc900Italic = &excelize.Font{
		Color:  "#18181b",
		Italic: true,
	}

	FontZinc950 = &excelize.Font{
		Color: "#09090b",
	}

	FontZinc950Italic = &excelize.Font{
		Color:  "#09090b",
		Italic: true,
	}

	FontNeutral50 = &excelize.Font{
		Color: "#fafafa",
	}

	FontNeutral50Italic = &excelize.Font{
		Color:  "#fafafa",
		Italic: true,
	}

	FontNeutral100 = &excelize.Font{
		Color: "#f5f5f5",
	}

	FontNeutral100Italic = &excelize.Font{
		Color:  "#f5f5f5",
		Italic: true,
	}

	FontNeutral200 = &excelize.Font{
		Color: "#e5e5e5",
	}

	FontNeutral200Italic = &excelize.Font{
		Color:  "#e5e5e5",
		Italic: true,
	}

	FontNeutral300 = &excelize.Font{
		Color: "#d4d4d4",
	}

	FontNeutral300Italic = &excelize.Font{
		Color:  "#d4d4d4",
		Italic: true,
	}

	FontNeutral400 = &excelize.Font{
		Color: "#a3a3a3",
	}

	FontNeutral400Italic = &excelize.Font{
		Color:  "#a3a3a3",
		Italic: true,
	}

	FontNeutral500 = &excelize.Font{
		Color: "#737373",
	}

	FontNeutral500Italic = &excelize.Font{
		Color:  "#737373",
		Italic: true,
	}

	FontNeutral600 = &excelize.Font{
		Color: "#525252",
	}

	FontNeutral600Italic = &excelize.Font{
		Color:  "#525252",
		Italic: true,
	}

	FontNeutral700 = &excelize.Font{
		Color: "#404040",
	}

	FontNeutral700Italic = &excelize.Font{
		Color:  "#404040",
		Italic: true,
	}

	FontNeutral800 = &excelize.Font{
		Color: "#262626",
	}

	FontNeutral800Italic = &excelize.Font{
		Color:  "#262626",
		Italic: true,
	}

	FontNeutral900 = &excelize.Font{
		Color: "#171717",
	}

	FontNeutral900Italic = &excelize.Font{
		Color:  "#171717",
		Italic: true,
	}

	FontNeutral950 = &excelize.Font{
		Color: "#0a0a0a",
	}

	FontNeutral950Italic = &excelize.Font{
		Color:  "#0a0a0a",
		Italic: true,
	}

	FontStone50 = &excelize.Font{
		Color: "#fafaf9",
	}

	FontStone50Italic = &excelize.Font{
		Color:  "#fafaf9",
		Italic: true,
	}

	FontStone100 = &excelize.Font{
		Color: "#f5f5f4",
	}

	FontStone100Italic = &excelize.Font{
		Color:  "#f5f5f4",
		Italic: true,
	}

	FontStone200 = &excelize.Font{
		Color: "#e7e5e4",
	}

	FontStone200Italic = &excelize.Font{
		Color:  "#e7e5e4",
		Italic: true,
	}

	FontStone300 = &excelize.Font{
		Color: "#d6d3d1",
	}

	FontStone300Italic = &excelize.Font{
		Color:  "#d6d3d1",
		Italic: true,
	}

	FontStone400 = &excelize.Font{
		Color: "#a8a29e",
	}

	FontStone400Italic = &excelize.Font{
		Color:  "#a8a29e",
		Italic: true,
	}

	FontStone500 = &excelize.Font{
		Color: "#78716c",
	}

	FontStone500Italic = &excelize.Font{
		Color:  "#78716c",
		Italic: true,
	}

	FontStone600 = &excelize.Font{
		Color: "#57534e",
	}

	FontStone600Italic = &excelize.Font{
		Color:  "#57534e",
		Italic: true,
	}

	FontStone700 = &excelize.Font{
		Color: "#44403c",
	}

	FontStone700Italic = &excelize.Font{
		Color:  "#44403c",
		Italic: true,
	}

	FontStone800 = &excelize.Font{
		Color: "#292524",
	}

	FontStone800Italic = &excelize.Font{
		Color:  "#292524",
		Italic: true,
	}

	FontStone900 = &excelize.Font{
		Color: "#1c1917",
	}

	FontStone900Italic = &excelize.Font{
		Color:  "#1c1917",
		Italic: true,
	}

	FontStone950 = &excelize.Font{
		Color: "#0c0a09",
	}

	FontStone950Italic = &excelize.Font{
		Color:  "#0c0a09",
		Italic: true,
	}

	FontRed50 = &excelize.Font{
		Color: "#fef2f2",
	}

	FontRed50Italic = &excelize.Font{
		Color:  "#fef2f2",
		Italic: true,
	}

	FontRed100 = &excelize.Font{
		Color: "#fee2e2",
	}

	FontRed100Italic = &excelize.Font{
		Color:  "#fee2e2",
		Italic: true,
	}

	FontRed200 = &excelize.Font{
		Color: "#fecaca",
	}

	FontRed200Italic = &excelize.Font{
		Color:  "#fecaca",
		Italic: true,
	}

	FontRed300 = &excelize.Font{
		Color: "#fca5a5",
	}

	FontRed300Italic = &excelize.Font{
		Color:  "#fca5a5",
		Italic: true,
	}

	FontRed400 = &excelize.Font{
		Color: "#f87171",
	}

	FontRed400Italic = &excelize.Font{
		Color:  "#f87171",
		Italic: true,
	}

	FontRed500 = &excelize.Font{
		Color: "#ef4444",
	}

	FontRed500Italic = &excelize.Font{
		Color:  "#ef4444",
		Italic: true,
	}

	FontRed600 = &excelize.Font{
		Color: "#dc2626",
	}

	FontRed600Italic = &excelize.Font{
		Color:  "#dc2626",
		Italic: true,
	}

	FontRed700 = &excelize.Font{
		Color: "#b91c1c",
	}

	FontRed700Italic = &excelize.Font{
		Color:  "#b91c1c",
		Italic: true,
	}

	FontRed800 = &excelize.Font{
		Color: "#991b1b",
	}

	FontRed800Italic = &excelize.Font{
		Color:  "#991b1b",
		Italic: true,
	}

	FontRed900 = &excelize.Font{
		Color: "#7f1d1d",
	}

	FontRed900Italic = &excelize.Font{
		Color:  "#7f1d1d",
		Italic: true,
	}

	FontRed950 = &excelize.Font{
		Color: "#450a0a",
	}

	FontRed950Italic = &excelize.Font{
		Color:  "#450a0a",
		Italic: true,
	}

	FontOrange50 = &excelize.Font{
		Color: "#fff7ed",
	}

	FontOrange50Italic = &excelize.Font{
		Color:  "#fff7ed",
		Italic: true,
	}

	FontOrange100 = &excelize.Font{
		Color: "#ffedd5",
	}

	FontOrange100Italic = &excelize.Font{
		Color:  "#ffedd5",
		Italic: true,
	}

	FontOrange200 = &excelize.Font{
		Color: "#fed7aa",
	}

	FontOrange200Italic = &excelize.Font{
		Color:  "#fed7aa",
		Italic: true,
	}

	FontOrange300 = &excelize.Font{
		Color: "#fdba74",
	}

	FontOrange300Italic = &excelize.Font{
		Color:  "#fdba74",
		Italic: true,
	}

	FontOrange400 = &excelize.Font{
		Color: "#fb923c",
	}

	FontOrange400Italic = &excelize.Font{
		Color:  "#fb923c",
		Italic: true,
	}

	FontOrange500 = &excelize.Font{
		Color: "#f97316",
	}

	FontOrange500Italic = &excelize.Font{
		Color:  "#f97316",
		Italic: true,
	}

	FontOrange600 = &excelize.Font{
		Color: "#ea580c",
	}

	FontOrange600Italic = &excelize.Font{
		Color:  "#ea580c",
		Italic: true,
	}

	FontOrange700 = &excelize.Font{
		Color: "#c2410c",
	}

	FontOrange700Italic = &excelize.Font{
		Color:  "#c2410c",
		Italic: true,
	}

	FontOrange800 = &excelize.Font{
		Color: "#9a3412",
	}

	FontOrange800Italic = &excelize.Font{
		Color:  "#9a3412",
		Italic: true,
	}

	FontOrange900 = &excelize.Font{
		Color: "#7c2d12",
	}

	FontOrange900Italic = &excelize.Font{
		Color:  "#7c2d12",
		Italic: true,
	}

	FontOrange950 = &excelize.Font{
		Color: "#431407",
	}

	FontOrange950Italic = &excelize.Font{
		Color:  "#431407",
		Italic: true,
	}

	FontAmber50 = &excelize.Font{
		Color: "#fffbeb",
	}

	FontAmber50Italic = &excelize.Font{
		Color:  "#fffbeb",
		Italic: true,
	}

	FontAmber100 = &excelize.Font{
		Color: "#fef3c7",
	}

	FontAmber100Italic = &excelize.Font{
		Color:  "#fef3c7",
		Italic: true,
	}

	FontAmber200 = &excelize.Font{
		Color: "#fde68a",
	}

	FontAmber200Italic = &excelize.Font{
		Color:  "#fde68a",
		Italic: true,
	}

	FontAmber300 = &excelize.Font{
		Color: "#fcd34d",
	}

	FontAmber300Italic = &excelize.Font{
		Color:  "#fcd34d",
		Italic: true,
	}

	FontAmber400 = &excelize.Font{
		Color: "#fbbf24",
	}

	FontAmber400Italic = &excelize.Font{
		Color:  "#fbbf24",
		Italic: true,
	}

	FontAmber500 = &excelize.Font{
		Color: "#f59e0b",
	}

	FontAmber500Italic = &excelize.Font{
		Color:  "#f59e0b",
		Italic: true,
	}

	FontAmber600 = &excelize.Font{
		Color: "#d97706",
	}

	FontAmber600Italic = &excelize.Font{
		Color:  "#d97706",
		Italic: true,
	}

	FontAmber700 = &excelize.Font{
		Color: "#b45309",
	}

	FontAmber700Italic = &excelize.Font{
		Color:  "#b45309",
		Italic: true,
	}

	FontAmber800 = &excelize.Font{
		Color: "#92400e",
	}

	FontAmber800Italic = &excelize.Font{
		Color:  "#92400e",
		Italic: true,
	}

	FontAmber900 = &excelize.Font{
		Color: "#78350f",
	}

	FontAmber900Italic = &excelize.Font{
		Color:  "#78350f",
		Italic: true,
	}

	FontAmber950 = &excelize.Font{
		Color: "#451a03",
	}

	FontAmber950Italic = &excelize.Font{
		Color:  "#451a03",
		Italic: true,
	}

	FontYellow50 = &excelize.Font{
		Color: "#fefce8",
	}

	FontYellow50Italic = &excelize.Font{
		Color:  "#fefce8",
		Italic: true,
	}

	FontYellow100 = &excelize.Font{
		Color: "#fef9c3",
	}

	FontYellow100Italic = &excelize.Font{
		Color:  "#fef9c3",
		Italic: true,
	}

	FontYellow200 = &excelize.Font{
		Color: "#fef08a",
	}

	FontYellow200Italic = &excelize.Font{
		Color:  "#fef08a",
		Italic: true,
	}

	FontYellow300 = &excelize.Font{
		Color: "#fde047",
	}

	FontYellow300Italic = &excelize.Font{
		Color:  "#fde047",
		Italic: true,
	}

	FontYellow400 = &excelize.Font{
		Color: "#facc15",
	}

	FontYellow400Italic = &excelize.Font{
		Color:  "#facc15",
		Italic: true,
	}

	FontYellow500 = &excelize.Font{
		Color: "#eab308",
	}

	FontYellow500Italic = &excelize.Font{
		Color:  "#eab308",
		Italic: true,
	}

	FontYellow600 = &excelize.Font{
		Color: "#ca8a04",
	}

	FontYellow600Italic = &excelize.Font{
		Color:  "#ca8a04",
		Italic: true,
	}

	FontYellow700 = &excelize.Font{
		Color: "#a16207",
	}

	FontYellow700Italic = &excelize.Font{
		Color:  "#a16207",
		Italic: true,
	}

	FontYellow800 = &excelize.Font{
		Color: "#854d0e",
	}

	FontYellow800Italic = &excelize.Font{
		Color:  "#854d0e",
		Italic: true,
	}

	FontYellow900 = &excelize.Font{
		Color: "#713f12",
	}

	FontYellow900Italic = &excelize.Font{
		Color:  "#713f12",
		Italic: true,
	}

	FontYellow950 = &excelize.Font{
		Color: "#422006",
	}

	FontYellow950Italic = &excelize.Font{
		Color:  "#422006",
		Italic: true,
	}

	FontLime50 = &excelize.Font{
		Color: "#f7fee7",
	}

	FontLime50Italic = &excelize.Font{
		Color:  "#f7fee7",
		Italic: true,
	}

	FontLime100 = &excelize.Font{
		Color: "#ecfccb",
	}

	FontLime100Italic = &excelize.Font{
		Color:  "#ecfccb",
		Italic: true,
	}

	FontLime200 = &excelize.Font{
		Color: "#d9f99d",
	}

	FontLime200Italic = &excelize.Font{
		Color:  "#d9f99d",
		Italic: true,
	}

	FontLime300 = &excelize.Font{
		Color: "#bef264",
	}

	FontLime300Italic = &excelize.Font{
		Color:  "#bef264",
		Italic: true,
	}

	FontLime400 = &excelize.Font{
		Color: "#a3e635",
	}

	FontLime400Italic = &excelize.Font{
		Color:  "#a3e635",
		Italic: true,
	}

	FontLime500 = &excelize.Font{
		Color: "#84cc16",
	}

	FontLime500Italic = &excelize.Font{
		Color:  "#84cc16",
		Italic: true,
	}

	FontLime600 = &excelize.Font{
		Color: "#65a30d",
	}

	FontLime600Italic = &excelize.Font{
		Color:  "#65a30d",
		Italic: true,
	}

	FontLime700 = &excelize.Font{
		Color: "#4d7c0f",
	}

	FontLime700Italic = &excelize.Font{
		Color:  "#4d7c0f",
		Italic: true,
	}

	FontLime800 = &excelize.Font{
		Color: "#3f6212",
	}

	FontLime800Italic = &excelize.Font{
		Color:  "#3f6212",
		Italic: true,
	}

	FontLime900 = &excelize.Font{
		Color: "#365314",
	}

	FontLime900Italic = &excelize.Font{
		Color:  "#365314",
		Italic: true,
	}

	FontLime950 = &excelize.Font{
		Color: "#1a2e05",
	}

	FontLime950Italic = &excelize.Font{
		Color:  "#1a2e05",
		Italic: true,
	}

	FontGreen50 = &excelize.Font{
		Color: "#f0fdf4",
	}

	FontGreen50Italic = &excelize.Font{
		Color:  "#f0fdf4",
		Italic: true,
	}

	FontGreen100 = &excelize.Font{
		Color: "#dcfce7",
	}

	FontGreen100Italic = &excelize.Font{
		Color:  "#dcfce7",
		Italic: true,
	}

	FontGreen200 = &excelize.Font{
		Color: "#bbf7d0",
	}

	FontGreen200Italic = &excelize.Font{
		Color:  "#bbf7d0",
		Italic: true,
	}

	FontGreen300 = &excelize.Font{
		Color: "#86efac",
	}

	FontGreen300Italic = &excelize.Font{
		Color:  "#86efac",
		Italic: true,
	}

	FontGreen400 = &excelize.Font{
		Color: "#4ade80",
	}

	FontGreen400Italic = &excelize.Font{
		Color:  "#4ade80",
		Italic: true,
	}

	FontGreen500 = &excelize.Font{
		Color: "#22c55e",
	}

	FontGreen500Italic = &excelize.Font{
		Color:  "#22c55e",
		Italic: true,
	}

	FontGreen600 = &excelize.Font{
		Color: "#16a34a",
	}

	FontGreen600Italic = &excelize.Font{
		Color:  "#16a34a",
		Italic: true,
	}

	FontGreen700 = &excelize.Font{
		Color: "#15803d",
	}

	FontGreen700Italic = &excelize.Font{
		Color:  "#15803d",
		Italic: true,
	}

	FontGreen800 = &excelize.Font{
		Color: "#166534",
	}

	FontGreen800Italic = &excelize.Font{
		Color:  "#166534",
		Italic: true,
	}

	FontGreen900 = &excelize.Font{
		Color: "#14532d",
	}

	FontGreen900Italic = &excelize.Font{
		Color:  "#14532d",
		Italic: true,
	}

	FontGreen950 = &excelize.Font{
		Color: "#052e16",
	}

	FontGreen950Italic = &excelize.Font{
		Color:  "#052e16",
		Italic: true,
	}

	FontEmerald50 = &excelize.Font{
		Color: "#ecfdf5",
	}

	FontEmerald50Italic = &excelize.Font{
		Color:  "#ecfdf5",
		Italic: true,
	}

	FontEmerald100 = &excelize.Font{
		Color: "#d1fae5",
	}

	FontEmerald100Italic = &excelize.Font{
		Color:  "#d1fae5",
		Italic: true,
	}

	FontEmerald200 = &excelize.Font{
		Color: "#a7f3d0",
	}

	FontEmerald200Italic = &excelize.Font{
		Color:  "#a7f3d0",
		Italic: true,
	}

	FontEmerald300 = &excelize.Font{
		Color: "#6ee7b7",
	}

	FontEmerald300Italic = &excelize.Font{
		Color:  "#6ee7b7",
		Italic: true,
	}

	FontEmerald400 = &excelize.Font{
		Color: "#34d399",
	}

	FontEmerald400Italic = &excelize.Font{
		Color:  "#34d399",
		Italic: true,
	}

	FontEmerald500 = &excelize.Font{
		Color: "#10b981",
	}

	FontEmerald500Italic = &excelize.Font{
		Color:  "#10b981",
		Italic: true,
	}

	FontEmerald600 = &excelize.Font{
		Color: "#059669",
	}

	FontEmerald600Italic = &excelize.Font{
		Color:  "#059669",
		Italic: true,
	}

	FontEmerald700 = &excelize.Font{
		Color: "#047857",
	}

	FontEmerald700Italic = &excelize.Font{
		Color:  "#047857",
		Italic: true,
	}

	FontEmerald800 = &excelize.Font{
		Color: "#065f46",
	}

	FontEmerald800Italic = &excelize.Font{
		Color:  "#065f46",
		Italic: true,
	}

	FontEmerald900 = &excelize.Font{
		Color: "#064e3b",
	}

	FontEmerald900Italic = &excelize.Font{
		Color:  "#064e3b",
		Italic: true,
	}

	FontEmerald950 = &excelize.Font{
		Color: "#022c22",
	}

	FontEmerald950Italic = &excelize.Font{
		Color:  "#022c22",
		Italic: true,
	}

	FontTeal50 = &excelize.Font{
		Color: "#f0fdfa",
	}

	FontTeal50Italic = &excelize.Font{
		Color:  "#f0fdfa",
		Italic: true,
	}

	FontTeal100 = &excelize.Font{
		Color: "#ccfbf1",
	}

	FontTeal100Italic = &excelize.Font{
		Color:  "#ccfbf1",
		Italic: true,
	}

	FontTeal200 = &excelize.Font{
		Color: "#99f6e4",
	}

	FontTeal200Italic = &excelize.Font{
		Color:  "#99f6e4",
		Italic: true,
	}

	FontTeal300 = &excelize.Font{
		Color: "#5eead4",
	}

	FontTeal300Italic = &excelize.Font{
		Color:  "#5eead4",
		Italic: true,
	}

	FontTeal400 = &excelize.Font{
		Color: "#2dd4bf",
	}

	FontTeal400Italic = &excelize.Font{
		Color:  "#2dd4bf",
		Italic: true,
	}

	FontTeal500 = &excelize.Font{
		Color: "#14b8a6",
	}

	FontTeal500Italic = &excelize.Font{
		Color:  "#14b8a6",
		Italic: true,
	}

	FontTeal600 = &excelize.Font{
		Color: "#0d9488",
	}

	FontTeal600Italic = &excelize.Font{
		Color:  "#0d9488",
		Italic: true,
	}

	FontTeal700 = &excelize.Font{
		Color: "#0f766e",
	}

	FontTeal700Italic = &excelize.Font{
		Color:  "#0f766e",
		Italic: true,
	}

	FontTeal800 = &excelize.Font{
		Color: "#115e59",
	}

	FontTeal800Italic = &excelize.Font{
		Color:  "#115e59",
		Italic: true,
	}

	FontTeal900 = &excelize.Font{
		Color: "#134e4a",
	}

	FontTeal900Italic = &excelize.Font{
		Color:  "#134e4a",
		Italic: true,
	}

	FontTeal950 = &excelize.Font{
		Color: "#042f2e",
	}

	FontTeal950Italic = &excelize.Font{
		Color:  "#042f2e",
		Italic: true,
	}

	FontCyan50 = &excelize.Font{
		Color: "#ecfeff",
	}

	FontCyan50Italic = &excelize.Font{
		Color:  "#ecfeff",
		Italic: true,
	}

	FontCyan100 = &excelize.Font{
		Color: "#cffafe",
	}

	FontCyan100Italic = &excelize.Font{
		Color:  "#cffafe",
		Italic: true,
	}

	FontCyan200 = &excelize.Font{
		Color: "#a5f3fc",
	}

	FontCyan200Italic = &excelize.Font{
		Color:  "#a5f3fc",
		Italic: true,
	}

	FontCyan300 = &excelize.Font{
		Color: "#67e8f9",
	}

	FontCyan300Italic = &excelize.Font{
		Color:  "#67e8f9",
		Italic: true,
	}

	FontCyan400 = &excelize.Font{
		Color: "#22d3ee",
	}

	FontCyan400Italic = &excelize.Font{
		Color:  "#22d3ee",
		Italic: true,
	}

	FontCyan500 = &excelize.Font{
		Color: "#06b6d4",
	}

	FontCyan500Italic = &excelize.Font{
		Color:  "#06b6d4",
		Italic: true,
	}

	FontCyan600 = &excelize.Font{
		Color: "#0891b2",
	}

	FontCyan600Italic = &excelize.Font{
		Color:  "#0891b2",
		Italic: true,
	}

	FontCyan700 = &excelize.Font{
		Color: "#0e7490",
	}

	FontCyan700Italic = &excelize.Font{
		Color:  "#0e7490",
		Italic: true,
	}

	FontCyan800 = &excelize.Font{
		Color: "#155e75",
	}

	FontCyan800Italic = &excelize.Font{
		Color:  "#155e75",
		Italic: true,
	}

	FontCyan900 = &excelize.Font{
		Color: "#164e63",
	}

	FontCyan900Italic = &excelize.Font{
		Color:  "#164e63",
		Italic: true,
	}

	FontCyan950 = &excelize.Font{
		Color: "#083344",
	}

	FontCyan950Italic = &excelize.Font{
		Color:  "#083344",
		Italic: true,
	}

	FontSky50 = &excelize.Font{
		Color: "#f0f9ff",
	}

	FontSky50Italic = &excelize.Font{
		Color:  "#f0f9ff",
		Italic: true,
	}

	FontSky100 = &excelize.Font{
		Color: "#e0f2fe",
	}

	FontSky100Italic = &excelize.Font{
		Color:  "#e0f2fe",
		Italic: true,
	}

	FontSky200 = &excelize.Font{
		Color: "#bae6fd",
	}

	FontSky200Italic = &excelize.Font{
		Color:  "#bae6fd",
		Italic: true,
	}

	FontSky300 = &excelize.Font{
		Color: "#7dd3fc",
	}

	FontSky300Italic = &excelize.Font{
		Color:  "#7dd3fc",
		Italic: true,
	}

	FontSky400 = &excelize.Font{
		Color: "#38bdf8",
	}

	FontSky400Italic = &excelize.Font{
		Color:  "#38bdf8",
		Italic: true,
	}

	FontSky500 = &excelize.Font{
		Color: "#0ea5e9",
	}

	FontSky500Italic = &excelize.Font{
		Color:  "#0ea5e9",
		Italic: true,
	}

	FontSky600 = &excelize.Font{
		Color: "#0284c7",
	}

	FontSky600Italic = &excelize.Font{
		Color:  "#0284c7",
		Italic: true,
	}

	FontSky700 = &excelize.Font{
		Color: "#0369a1",
	}

	FontSky700Italic = &excelize.Font{
		Color:  "#0369a1",
		Italic: true,
	}

	FontSky800 = &excelize.Font{
		Color: "#075985",
	}

	FontSky800Italic = &excelize.Font{
		Color:  "#075985",
		Italic: true,
	}

	FontSky900 = &excelize.Font{
		Color: "#0c4a6e",
	}

	FontSky900Italic = &excelize.Font{
		Color:  "#0c4a6e",
		Italic: true,
	}

	FontSky950 = &excelize.Font{
		Color: "#082f49",
	}

	FontSky950Italic = &excelize.Font{
		Color:  "#082f49",
		Italic: true,
	}

	FontBlue50 = &excelize.Font{
		Color: "#eff6ff",
	}

	FontBlue50Italic = &excelize.Font{
		Color:  "#eff6ff",
		Italic: true,
	}

	FontBlue100 = &excelize.Font{
		Color: "#dbeafe",
	}

	FontBlue100Italic = &excelize.Font{
		Color:  "#dbeafe",
		Italic: true,
	}

	FontBlue200 = &excelize.Font{
		Color: "#bfdbfe",
	}

	FontBlue200Italic = &excelize.Font{
		Color:  "#bfdbfe",
		Italic: true,
	}

	FontBlue300 = &excelize.Font{
		Color: "#93c5fd",
	}

	FontBlue300Italic = &excelize.Font{
		Color:  "#93c5fd",
		Italic: true,
	}

	FontBlue400 = &excelize.Font{
		Color: "#60a5fa",
	}

	FontBlue400Italic = &excelize.Font{
		Color:  "#60a5fa",
		Italic: true,
	}

	FontBlue500 = &excelize.Font{
		Color: "#3b82f6",
	}

	FontBlue500Italic = &excelize.Font{
		Color:  "#3b82f6",
		Italic: true,
	}

	FontBlue600 = &excelize.Font{
		Color: "#2563eb",
	}

	FontBlue600Italic = &excelize.Font{
		Color:  "#2563eb",
		Italic: true,
	}

	FontBlue700 = &excelize.Font{
		Color: "#1d4ed8",
	}

	FontBlue700Italic = &excelize.Font{
		Color:  "#1d4ed8",
		Italic: true,
	}

	FontBlue800 = &excelize.Font{
		Color: "#1e40af",
	}

	FontBlue800Italic = &excelize.Font{
		Color:  "#1e40af",
		Italic: true,
	}

	FontBlue900 = &excelize.Font{
		Color: "#1e3a8a",
	}

	FontBlue900Italic = &excelize.Font{
		Color:  "#1e3a8a",
		Italic: true,
	}

	FontBlue950 = &excelize.Font{
		Color: "#172554",
	}

	FontBlue950Italic = &excelize.Font{
		Color:  "#172554",
		Italic: true,
	}

	FontIndigo50 = &excelize.Font{
		Color: "#eef2ff",
	}

	FontIndigo50Italic = &excelize.Font{
		Color:  "#eef2ff",
		Italic: true,
	}

	FontIndigo100 = &excelize.Font{
		Color: "#e0e7ff",
	}

	FontIndigo100Italic = &excelize.Font{
		Color:  "#e0e7ff",
		Italic: true,
	}

	FontIndigo200 = &excelize.Font{
		Color: "#c7d2fe",
	}

	FontIndigo200Italic = &excelize.Font{
		Color:  "#c7d2fe",
		Italic: true,
	}

	FontIndigo300 = &excelize.Font{
		Color: "#a5b4fc",
	}

	FontIndigo300Italic = &excelize.Font{
		Color:  "#a5b4fc",
		Italic: true,
	}

	FontIndigo400 = &excelize.Font{
		Color: "#818cf8",
	}

	FontIndigo400Italic = &excelize.Font{
		Color:  "#818cf8",
		Italic: true,
	}

	FontIndigo500 = &excelize.Font{
		Color: "#6366f1",
	}

	FontIndigo500Italic = &excelize.Font{
		Color:  "#6366f1",
		Italic: true,
	}

	FontIndigo600 = &excelize.Font{
		Color: "#4f46e5",
	}

	FontIndigo600Italic = &excelize.Font{
		Color:  "#4f46e5",
		Italic: true,
	}

	FontIndigo700 = &excelize.Font{
		Color: "#4338ca",
	}

	FontIndigo700Italic = &excelize.Font{
		Color:  "#4338ca",
		Italic: true,
	}

	FontIndigo800 = &excelize.Font{
		Color: "#3730a3",
	}

	FontIndigo800Italic = &excelize.Font{
		Color:  "#3730a3",
		Italic: true,
	}

	FontIndigo900 = &excelize.Font{
		Color: "#312e81",
	}

	FontIndigo900Italic = &excelize.Font{
		Color:  "#312e81",
		Italic: true,
	}

	FontIndigo950 = &excelize.Font{
		Color: "#1e1b4b",
	}

	FontIndigo950Italic = &excelize.Font{
		Color:  "#1e1b4b",
		Italic: true,
	}

	FontViolet50 = &excelize.Font{
		Color: "#f5f3ff",
	}

	FontViolet50Italic = &excelize.Font{
		Color:  "#f5f3ff",
		Italic: true,
	}

	FontViolet100 = &excelize.Font{
		Color: "#ede9fe",
	}

	FontViolet100Italic = &excelize.Font{
		Color:  "#ede9fe",
		Italic: true,
	}

	FontViolet200 = &excelize.Font{
		Color: "#ddd6fe",
	}

	FontViolet200Italic = &excelize.Font{
		Color:  "#ddd6fe",
		Italic: true,
	}

	FontViolet300 = &excelize.Font{
		Color: "#c4b5fd",
	}

	FontViolet300Italic = &excelize.Font{
		Color:  "#c4b5fd",
		Italic: true,
	}

	FontViolet400 = &excelize.Font{
		Color: "#a78bfa",
	}

	FontViolet400Italic = &excelize.Font{
		Color:  "#a78bfa",
		Italic: true,
	}

	FontViolet500 = &excelize.Font{
		Color: "#8b5cf6",
	}

	FontViolet500Italic = &excelize.Font{
		Color:  "#8b5cf6",
		Italic: true,
	}

	FontViolet600 = &excelize.Font{
		Color: "#7c3aed",
	}

	FontViolet600Italic = &excelize.Font{
		Color:  "#7c3aed",
		Italic: true,
	}

	FontViolet700 = &excelize.Font{
		Color: "#6d28d9",
	}

	FontViolet700Italic = &excelize.Font{
		Color:  "#6d28d9",
		Italic: true,
	}

	FontViolet800 = &excelize.Font{
		Color: "#5b21b6",
	}

	FontViolet800Italic = &excelize.Font{
		Color:  "#5b21b6",
		Italic: true,
	}

	FontViolet900 = &excelize.Font{
		Color: "#4c1d95",
	}

	FontViolet900Italic = &excelize.Font{
		Color:  "#4c1d95",
		Italic: true,
	}

	FontViolet950 = &excelize.Font{
		Color: "#2e1065",
	}

	FontViolet950Italic = &excelize.Font{
		Color:  "#2e1065",
		Italic: true,
	}

	FontPurple50 = &excelize.Font{
		Color: "#faf5ff",
	}

	FontPurple50Italic = &excelize.Font{
		Color:  "#faf5ff",
		Italic: true,
	}

	FontPurple100 = &excelize.Font{
		Color: "#f3e8ff",
	}

	FontPurple100Italic = &excelize.Font{
		Color:  "#f3e8ff",
		Italic: true,
	}

	FontPurple200 = &excelize.Font{
		Color: "#e9d5ff",
	}

	FontPurple200Italic = &excelize.Font{
		Color:  "#e9d5ff",
		Italic: true,
	}

	FontPurple300 = &excelize.Font{
		Color: "#d8b4fe",
	}

	FontPurple300Italic = &excelize.Font{
		Color:  "#d8b4fe",
		Italic: true,
	}

	FontPurple400 = &excelize.Font{
		Color: "#c084fc",
	}

	FontPurple400Italic = &excelize.Font{
		Color:  "#c084fc",
		Italic: true,
	}

	FontPurple500 = &excelize.Font{
		Color: "#a855f7",
	}

	FontPurple500Italic = &excelize.Font{
		Color:  "#a855f7",
		Italic: true,
	}

	FontPurple600 = &excelize.Font{
		Color: "#9333ea",
	}

	FontPurple600Italic = &excelize.Font{
		Color:  "#9333ea",
		Italic: true,
	}

	FontPurple700 = &excelize.Font{
		Color: "#7e22ce",
	}

	FontPurple700Italic = &excelize.Font{
		Color:  "#7e22ce",
		Italic: true,
	}

	FontPurple800 = &excelize.Font{
		Color: "#6b21a8",
	}

	FontPurple800Italic = &excelize.Font{
		Color:  "#6b21a8",
		Italic: true,
	}

	FontPurple900 = &excelize.Font{
		Color: "#581c87",
	}

	FontPurple900Italic = &excelize.Font{
		Color:  "#581c87",
		Italic: true,
	}

	FontPurple950 = &excelize.Font{
		Color: "#3b0764",
	}

	FontPurple950Italic = &excelize.Font{
		Color:  "#3b0764",
		Italic: true,
	}

	FontFuchsia50 = &excelize.Font{
		Color: "#fdf4ff",
	}

	FontFuchsia50Italic = &excelize.Font{
		Color:  "#fdf4ff",
		Italic: true,
	}

	FontFuchsia100 = &excelize.Font{
		Color: "#fae8ff",
	}

	FontFuchsia100Italic = &excelize.Font{
		Color:  "#fae8ff",
		Italic: true,
	}

	FontFuchsia200 = &excelize.Font{
		Color: "#f5d0fe",
	}

	FontFuchsia200Italic = &excelize.Font{
		Color:  "#f5d0fe",
		Italic: true,
	}

	FontFuchsia300 = &excelize.Font{
		Color: "#f0abfc",
	}

	FontFuchsia300Italic = &excelize.Font{
		Color:  "#f0abfc",
		Italic: true,
	}

	FontFuchsia400 = &excelize.Font{
		Color: "#e879f9",
	}

	FontFuchsia400Italic = &excelize.Font{
		Color:  "#e879f9",
		Italic: true,
	}

	FontFuchsia500 = &excelize.Font{
		Color: "#d946ef",
	}

	FontFuchsia500Italic = &excelize.Font{
		Color:  "#d946ef",
		Italic: true,
	}

	FontFuchsia600 = &excelize.Font{
		Color: "#c026d3",
	}

	FontFuchsia600Italic = &excelize.Font{
		Color:  "#c026d3",
		Italic: true,
	}

	FontFuchsia700 = &excelize.Font{
		Color: "#a21caf",
	}

	FontFuchsia700Italic = &excelize.Font{
		Color:  "#a21caf",
		Italic: true,
	}

	FontFuchsia800 = &excelize.Font{
		Color: "#86198f",
	}

	FontFuchsia800Italic = &excelize.Font{
		Color:  "#86198f",
		Italic: true,
	}

	FontFuchsia900 = &excelize.Font{
		Color: "#701a75",
	}

	FontFuchsia900Italic = &excelize.Font{
		Color:  "#701a75",
		Italic: true,
	}

	FontFuchsia950 = &excelize.Font{
		Color: "#4a044e",
	}

	FontFuchsia950Italic = &excelize.Font{
		Color:  "#4a044e",
		Italic: true,
	}

	FontPink50 = &excelize.Font{
		Color: "#fdf2f8",
	}

	FontPink50Italic = &excelize.Font{
		Color:  "#fdf2f8",
		Italic: true,
	}

	FontPink100 = &excelize.Font{
		Color: "#fce7f3",
	}

	FontPink100Italic = &excelize.Font{
		Color:  "#fce7f3",
		Italic: true,
	}

	FontPink200 = &excelize.Font{
		Color: "#fbcfe8",
	}

	FontPink200Italic = &excelize.Font{
		Color:  "#fbcfe8",
		Italic: true,
	}

	FontPink300 = &excelize.Font{
		Color: "#f9a8d4",
	}

	FontPink300Italic = &excelize.Font{
		Color:  "#f9a8d4",
		Italic: true,
	}

	FontPink400 = &excelize.Font{
		Color: "#f472b6",
	}

	FontPink400Italic = &excelize.Font{
		Color:  "#f472b6",
		Italic: true,
	}

	FontPink500 = &excelize.Font{
		Color: "#ec4899",
	}

	FontPink500Italic = &excelize.Font{
		Color:  "#ec4899",
		Italic: true,
	}

	FontPink600 = &excelize.Font{
		Color: "#db2777",
	}

	FontPink600Italic = &excelize.Font{
		Color:  "#db2777",
		Italic: true,
	}

	FontPink700 = &excelize.Font{
		Color: "#be185d",
	}

	FontPink700Italic = &excelize.Font{
		Color:  "#be185d",
		Italic: true,
	}

	FontPink800 = &excelize.Font{
		Color: "#9d174d",
	}

	FontPink800Italic = &excelize.Font{
		Color:  "#9d174d",
		Italic: true,
	}

	FontPink900 = &excelize.Font{
		Color: "#831843",
	}

	FontPink900Italic = &excelize.Font{
		Color:  "#831843",
		Italic: true,
	}

	FontPink950 = &excelize.Font{
		Color: "#500724",
	}

	FontPink950Italic = &excelize.Font{
		Color:  "#500724",
		Italic: true,
	}

	FontRose50 = &excelize.Font{
		Color: "#fff1f2",
	}

	FontRose50Italic = &excelize.Font{
		Color:  "#fff1f2",
		Italic: true,
	}

	FontRose100 = &excelize.Font{
		Color: "#ffe4e6",
	}

	FontRose100Italic = &excelize.Font{
		Color:  "#ffe4e6",
		Italic: true,
	}

	FontRose200 = &excelize.Font{
		Color: "#fecdd3",
	}

	FontRose200Italic = &excelize.Font{
		Color:  "#fecdd3",
		Italic: true,
	}

	FontRose300 = &excelize.Font{
		Color: "#fda4af",
	}

	FontRose300Italic = &excelize.Font{
		Color:  "#fda4af",
		Italic: true,
	}

	FontRose400 = &excelize.Font{
		Color: "#fb7185",
	}

	FontRose400Italic = &excelize.Font{
		Color:  "#fb7185",
		Italic: true,
	}

	FontRose500 = &excelize.Font{
		Color: "#f43f5e",
	}

	FontRose500Italic = &excelize.Font{
		Color:  "#f43f5e",
		Italic: true,
	}

	FontRose600 = &excelize.Font{
		Color: "#e11d48",
	}

	FontRose600Italic = &excelize.Font{
		Color:  "#e11d48",
		Italic: true,
	}

	FontRose700 = &excelize.Font{
		Color: "#be123c",
	}

	FontRose700Italic = &excelize.Font{
		Color:  "#be123c",
		Italic: true,
	}

	FontRose800 = &excelize.Font{
		Color: "#9f1239",
	}

	FontRose800Italic = &excelize.Font{
		Color:  "#9f1239",
		Italic: true,
	}

	FontRose900 = &excelize.Font{
		Color: "#881337",
	}

	FontRose900Italic = &excelize.Font{
		Color:  "#881337",
		Italic: true,
	}

	FontRose950 = &excelize.Font{
		Color: "#4c0519",
	}

	FontRose950Italic = &excelize.Font{
		Color:  "#4c0519",
		Italic: true,
	}
)
