import os
import re
import glob
from collections import defaultdict

def find_files_with_name(start_path, filename):
    result_files = []
    
    for file_path in glob.iglob(os.path.join(start_path, "**", filename), recursive=True):
        result_files.append(file_path)
    
    return result_files

def find_module_name(s):
    start_index = s.find("/x/") + len("/x/")
    end_index = s.find("/", start_index)
    if start_index != -1 and end_index != -1:
        return s[start_index:end_index]
    return None

def get_moduleName(file_path):
    if os.path.exists(param_path):
        with open(file_path,'r') as file:
            for line in file:
                if "ModuleName = " in line:
                    # need to be func
                    parts = line.split("=")
                    if len(parts) == 2:
                        return parts[1].strip().strip('"')
                    else:
                        print(file_path)
                        print("error")
                        exit(1)
    return None

def read_file(file_path):
    const_dict = {}
    error_dict = {}
    with open(file_path,'r') as file:
        for line in file:
            if "=" in line:
                # get const
                if "sdkerrors.Register" not in line: 
                    cleaned_line = line.replace("const", "").strip()
                    parts = cleaned_line.split("=")
                    if len(parts) == 2:
                        identifier, value = parts[0].strip(), parts[1].strip().strip('"')
                        const_dict[identifier]=value
                    else:
                        print(file_path)
                        print("error")
                        exit(1)
                # get error information
                else:
                    parts = line.split("=", 1)
                    error_name, value = parts[0].strip(), parts[1].strip()
                    # error info is like as sdkerrors.Register(...)
                    pattern = r"sdkerrors\.Register\((.*)\)"
                    match = re.search(pattern, value)
                    if match:
                        contents = match.group(1)
                        parts = contents.split(",", 2)
                        
                        code_space=parts[0].strip().strip('"')
                        code = parts[1].strip().strip('"')
                        description=parts[2].strip().strip('"')

                        if code_space in const_dict.keys():
                            code_space=const_dict[code_space]
                        
                        error_info = (code_space, code, description)
                        error_dict[error_name]=error_info
                    else:
                        print(file_path)
                        print(line)
                        print("failed")
    return const_dict,error_dict

if __name__ == "__main__":
    current_path = os.path.dirname(__file__)
    target_path = os.path.join(current_path,'..','..','x')

    # get all errors.go in x folder
    error_file = "errors.go"
    file_paths = find_files_with_name(target_path, error_file)
    if not(file_paths): 
        print("Not find target files in x folder")
        exit(1)

    # get module name and bind with paths (one module may have multiple errors.go)
    module_with_paths = defaultdict(list)
    for file_path in file_paths:
        module_name=find_module_name(file_path)
        if not(module_name):
            print("error file path, failed on %s" % file_path)
            exit(1)
        module_with_paths[module_name].append(file_path)

    # execute
    with open(target_path+"/errors.md", "w") as file:
        # category        
        file.write('<!-- TOC -->\n')
        file.write('Category\n')
        column_template='  * [{name}](#{link})\n'
        for module_name, file_paths in module_with_paths.items():
            file.write(column_template.format(name=module_name.capitalize(), link=module_name))
        file.write('<!-- TOC -->\n')

        # errors in each module
        for module_name, file_paths in module_with_paths.items():
            
            # table header
            file.write('## %s\n' % module_name.capitalize())
            file.write('\n')
            file.write('|Error Name|Codespace|Code|Description|\n')
            file.write('|:-|:-|:-|:-|\n')

            # table contents
            error_template='|{error_name}|{code_space}|{code}|{description}|\n'
            for file_path in file_paths:
                const_dict, error_dict = read_file(file_path)

                param_path = file_path.replace("errors.go", "key.go")
                moduleName=get_moduleName(param_path)
                if not(moduleName):
                    param_path = file_path.replace("errors.go", "keys.go")
                    moduleName=get_moduleName(param_path)
                for error_name, error_info in error_dict.items():
                    if error_info[0]=="ModuleName":
                        code_space=moduleName
                    else:
                        code_space=error_info[0]
                    file.write(error_template.format(error_name=error_name, code_space=code_space,code = error_info[1],description=error_info[2]))