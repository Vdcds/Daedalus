package catalog

var Categories = []Category{
	{
		ID:          "cli",
		Name:        "CLI Tools",
		Description: "Core terminal utilities",

		Packages: []Package{
			{
				Name:        "Git",
				Description: "Distributed version control",
				BrewName:    "git",
			},
			{
				Name:        "Ripgrep",
				Description: "Fast recursive search",
				BrewName:    "ripgrep",
			},
			{
				Name:        "FZF",
				Description: "Command-line fuzzy finder",
				BrewName:    "fzf",
			},
			{
				Name:        "Bat",
				Description: "Modern cat replacement",
				BrewName:    "bat",
			},
			{
				Name:        "Eza",
				Description: "Modern ls replacement",
				BrewName:    "eza",
			},
		},
	},

	{
		ID:          "editors",
		Name:        "Editors",
		Description: "Code editors & IDEs",

		Packages: []Package{
			{
				Name:        "Neovim",
				Description: "Hyperextensible Vim",
				BrewName:    "neovim",
			},
			{
				Name:        "Zed",
				Description: "High-performance editor",
				BrewName:    "zed",
			},
			{
				Name:        "Visual Studio Code",
				Description: "Microsoft editor",
				BrewName:    "visual-studio-code",
			},
		},
	},

	{
		ID:          "browsers",
		Name:        "Browsers",
		Description: "Modern web browsers",

		Packages: []Package{
			{
				Name:        "Zen Browser",
				Description: "Firefox-based productivity browser",
				BrewName:    "zen-browser",
			},
			{
				Name:        "Firefox",
				Description: "Mozilla Firefox",
				BrewName:    "firefox",
			},
			{
				Name:        "Brave",
				Description: "Privacy-focused Chromium browser",
				BrewName:    "brave-browser",
			},
			{
				Name:        "Google Chrome",
				Description: "Google's web browser",
				BrewName:    "google-chrome",
			},
		},
	},

	{
		ID:          "development",
		Name:        "Development",
		Description: "Developer tooling",

		Packages: []Package{
			{
				Name:        "Docker",
				Description: "Container platform",
				BrewName:    "docker",
			},
			{
				Name:        "Go",
				Description: "Go programming language",
				BrewName:    "go",
			},
			{
				Name:        "Node.js",
				Description: "JavaScript runtime",
				BrewName:    "node",
			},
		},
	},

	{
		ID:          "utilities",
		Name:        "Utilities",
		Description: "Everyday applications",

		Packages: []Package{
			{
				Name:        "Rectangle",
				Description: "Window manager",
				BrewName:    "rectangle",
			},
			{
				Name:        "Raycast",
				Description: "Launcher and productivity tool",
				BrewName:    "raycast",
			},
			{
				Name:        "Stats",
				Description: "System monitor",
				BrewName:    "stats",
			},
		},
	},
}
