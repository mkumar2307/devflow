# devflow
A fast Go CLI that lets developers understand and work with codebases using AI, directly from the terminal.

🚀 Core Features (MVP)      

1️⃣ Codebase Understanding        

```bash
devflow explain ./pkg/auth
devflow ask "Where is JWT validated?"
devflow summarize-pr 142
devflow debug panic.log
```

Capabilities:        

Summarize folder/module        
Explain file purpose       
Detect entry points      
Identify dependencies       
Generate high-level architecture overview      

2️⃣ Intelligent Q&A       

```bash
devflow ask "Where is JWT validated?"
devflow ask "How does payment retry work?"
```

Under the hood:        
Code indexing        
Embeddings      
Context-aware retrieval (RAG)     
LLM reasoning      

3️⃣ AI Refactoring       
x
```bash
devflow refactor ./service/user.go --goal="make this more testable"        
```
Capabilities:         
Suggest improvements         
Extract interfaces          
Improve naming         
Reduce duplication          
Add error handling         

4️⃣ Test Generation         

```bash
devflow test ./internal/cache
```

Generates:        
Unit tests        
Table-driven tests (Go idiomatic)        
Mock suggestions       

5️⃣ Smart Debugging        

```bash
devflow debug "nil pointer in order service"       
```

Possible:         
Analyze stack traces       
Trace call chain         
Suggest root cause           
Highlight risky patterns        

🧠 Architecture Design       

CLI Layer (Go + Cobra)         
  
spf13 (Cobra framework)         
Viper for config       
Bubble Tea (optional TUI mode)         
