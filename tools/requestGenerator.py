import os

pythonSdkPath = "/Users/yiz96/foss/zmxy-sdk/python/zmop/request/"

# process python sdk

def camelCase2_Case(str):
    res = []
    str += 'A'
    cur = ""
    for c in str:
        if c >= 'A' and c <= 'Z':
            res += [cur.lower()]
            cur = ""
        cur += c
    return "_".join(res)

def WriteSetter(fileHandler, className, param, type):
    fileHandler.write("\nfunc (m *" + className + ") Set" + param[0].upper() + param[1:] + "(" + param + " " + type + ") {\n")
    fileHandler.write("m." + param + " = " + param + "\n")
    fileHandler.write("}\n\n")

def WriteGetter(fileHandler, className, param, type):
    fileHandler.write("\nfunc (m *" + className + ") Get" + param[0].upper() + param[1:] + "() " + type + " {\n")
    fileHandler.write("return m." + param + "\n")
    fileHandler.write("}\n\n")

files = os.listdir(pythonSdkPath)
for basedir in files:
    if basedir == "__init__.py":
        continue
    members = []
    comments = []

    className = basedir.split(".py")[0]
    f = open(pythonSdkPath + basedir)
    fout = open("../requests/" + className + ".go", "w")

    fout.write("package requests\n\nimport (\n\"github.com/delostik/go-zmxy/interfaces\"\n)\n\n")
    fout.write("type " + className + " struct {\n")

    for i in range(3):
        f.readline()

    methodName = f.readline().split(" ")[2]
    
    for i in range(7):
        f.readline()
    
    line = f.readline()
    while line != "\n":
        member = line.strip().split(".__")[1].split(" ")[0]
        if member == "type":
            member = "opType"
        members += [member]
        comments += [line.strip().split('#')[1]]
        line = f.readline()
    for idx, member in enumerate(members):
        fout.write(member + " string //" + comments[idx] + "\n")

    fout.write("\ninterfaces.ZhimaRequestParams\n}\n\n")

    fout.write("func (m *" + className + ") InitBizParams(" + ' ,'.join(members) + " string) {\n")
    fout.write("m.ApiParams = &map[string]string{}\nm.FileParams = &map[string]string{}\n")

    for member in members:
        if member != "file":
            fout.write("\n(*m.ApiParams)[\"" + camelCase2_Case(member) + "\"] = " + member + "\n")
        else:
            fout.write("\n(*m.FileParams)[\"" + camelCase2_Case(member) + "\"] = " + member + "\n")
        fout.write("m." + member + " = " + member + "\n")
    fout.write("}\n")

    # ApiMethodName
    fout.write("\nfunc (*" + className + ") GetApiMethodName() string {\nreturn \"" + methodName + "\"\n}\n")

    WriteGetter(fout, className, "ApiParams", "*map[string]string")
    WriteGetter(fout, className, "FileParams", "*map[string]string")

    WriteSetter(fout, className, "ApiVersion", "string")
    WriteGetter(fout, className, "ApiVersion", "string")

    WriteSetter(fout, className, "Scene", "string")
    WriteGetter(fout, className, "Scene", "string")

    WriteSetter(fout, className, "Channel", "string")
    WriteGetter(fout, className, "Channel", "string")

    WriteSetter(fout, className, "Platform", "string")
    WriteGetter(fout, className, "Platform", "string")

    WriteSetter(fout, className, "ExtParams", "string")
    WriteGetter(fout, className, "ExtParams", "string")

    f.close()
    fout.close()
    
    os.system("go fmt ../requests/" + className + ".go")