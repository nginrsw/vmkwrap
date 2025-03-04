# **VMKWrap** 🚀 

VMKWrap is an automatic transpiler that converts all **VMK** files in your project to **Lua**, executes them using **LuaJIT**,
and then converts them back to **VMK**. It enables seamless execution of VMK `code`, making it useful for any project
that relies on **LuaJIT**.

## **⚠ Requirements**  
VMKWrap requires **LuaJIT**. Make sure LuaJIT is installed on your computer; otherwise, this tool will not work.

## **📌 Features** 

- ✅ Automatically transpiles **VMK → Lua** and back
- ✅ Runs code using **LuaJIT**
- ✅ Standalone binary, can be used anywhere
- ✅ Simple one-command execution with `run.sh`

---

**Question:** Why not just use the VMK interpreter or standard Lua?

**Answer:** This tool is actually designed for `Lua` game engines, which commonly
rely on LuaJIT. If you need a tool for a different purpose or interpreter,
feel free to modify this project locally.

---

## **🚀 Installation & Setup** 

### **1️⃣ Clone the Repository**  

```sh
git clone https://github.com/nginrsw/vmkwrap.git
cd vmkwrap
```

### **2️⃣ Build the Project**  

Run the following command to build VMKWrap which **stored** in `bin/`:

```sh
./run.sh
```

**vmkwrap** can be found in the bin/ folder.

### **3️⃣ Optional: Create symlink for Global Access**  

To make `vmkwrap` accessible from anywhere:

```sh
sudo ln -s ~/username/your-dir/bin/vmkwrap /usr/local/bin/
```

Then you can run it from any directory in your terminal:

```sh
vmkwrap main.vmk
```

---

## **📜 License** 

This project is licensed under the [MIT License](LICENSE).
