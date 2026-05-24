<div align="center">
  <img src="https://capsule-render.vercel.app/api?type=soft&color=5A5CD6&height=140&section=header&text=ColorMark&fontSize=52&fontColor=ffffff" width="100%" />
  
  <h3>Fast, beautiful, and lightweight Markdown reader for the terminal.</h3>
  
  <p>
    <img src="https://img.shields.io/github/license/mehmetalidsy/colormark?style=for-the-badge&color=blue" alt="License" />
    <img src="https://img.shields.io/github/stars/mehmetalidsy/colormark?style=for-the-badge&color=yellow" alt="Stars" />
    <img src="https://img.shields.io/github/forks/mehmetalidsy/colormark?style=for-the-badge&color=green" alt="Forks" />
    <img src="https://img.shields.io/badge/Made%20with-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  </p>
</div>

---

### 🚀 What is ColorMark?

ColorMark is a **blazing fast** Terminal User Interface (TUI) Markdown viewer built for developers who live in the terminal.

No more opening heavy editors or browsers just to read a `README.md` or documentation. ColorMark gives you a clean, beautiful, and highly responsive Markdown reading experience directly in your terminal.

**Why ColorMark?**
- Extremely lightweight and fast
- Beautiful rendering with your terminal theme
- Zero heavy runtime dependencies (single binary)
- Vim-like controls + mouse support

---

### Demo

<video src="https://github.com/mehmetalidsy/colormark/assets/12345678/12345678-1234-1234-1234-1234567890ab" width="100%" controls>Your browser does not support the video tag.</video>

---

### ✨ Features

- **Fast Rendering** with Glamour + custom optimizations
- **Vim-style Navigation** (`j/k`, `gg`, `G`, `Ctrl+D/U`)
- **Table of Contents** sidebar (coming soon)
- **Code block copying** with one key
- **Search** inside document
- **Multiple Themes** (dark/light + custom)
- **Mouse Support**
- **Configurable** via `~/.config/colormark/config.yaml`
- Single static binary (no Node.js, Python etc.)

---

### 📊 Comparison

| Feature              | ColorMark     | Glow          | mdcat     |
|----------------------|---------------|---------------|-----------|
| Speed                | Very Fast     | Fast          | Fast      |
| TUI Experience       | Rich          | Good          | Basic     |
| Customization        | High          | Medium        | Low       |
| Single Binary        | Yes           | Yes           | Yes       |
| Active Development   | Yes           | Yes           | Yes       |

---

### 🛠️ Installation

```bash
# Install directly
go install github.com/mehmetalidsy/colormark@latest

# Or from source
git clone https://github.com/mehmetalidsy/colormark.git
cd colormark
go install .

```

### Usage

```bash
colormark README.md
colormark docs/api.md
```

### 🎮 Key Bindings

| Key             | Action                    |
|-----------------|---------------------------|
| `j` / `↓`       | Scroll down               |
| `k` / `↑`       | Scroll up                 |
| `gg`            | Go to top                 |
| `G`             | Go to bottom              |
| `q` / `Ctrl+C`  | Quit                      |
| `c` (hover)     | Copy code block           |
| `/`             | Search                    |

## Roadmap

- [ ] Table of Contents sidebar
- [ ] Full text search
- [ ] Plugin support
- [ ] Export to PDF/HTML
- [ ] LSP integration (future)

### 🤝 Contributing

Contributions are welcome! Feel free to open issues or PRs.