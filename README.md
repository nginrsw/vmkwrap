# **VMKWrap** 🚀 

VMKWrap is an automatic transpiler that converts **VMK to Lua**, runs it using
**LuaJIT**, and then converts it back to **VMK**. It allows seamless execution
of VMK code, making it useful for anything using **LuaJIT**.

## **📌 Features** 

- ✅ Automatically transpiles **VMK → Lua** and back
- ✅ Runs Lua code using **LuaJIT**
- ✅ Standalone binary, can be used anywhere
- ✅ Simple one-command execution with `run.sh`

---

## **🚀 Installation & Setup** 

### **1️⃣ Clone the Repository**  

```sh
git clone https://github.com/nginrsw/vmkwrap.git
cd vmkwrap
```

### **2️⃣ Build the Project**  

Run the following command to build VMKWrap and **store it** in `bin/`:

```sh
./run.sh
```

**VMKWrap** can be found in the bin/ folder.

### **3️⃣ Optional: Move Binary for Global Access**  

To make `vmkwrap` accessible from anywhere:

```sh
sudo ln -s ~/username/your-dir/bin/vmkwrap /usr/local/bin/
```

Then you can run it from any directory:

```sh
vmkwrap main.vmk
```

---

## **📜 License** 

This project is licensed under the [MIT License](LICENSE).
