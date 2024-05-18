import json

def format_log_file(input_file, output_file):
    with open(input_file, 'r', encoding='utf-8') as infile, open(output_file, 'w', encoding='utf-8') as outfile:
        for line in infile:
            try:
                # 解析 JSON 数据
                log_entry = json.loads(line)
                # 写入格式化的 JSON 数据
                json.dump(log_entry, outfile, indent=4, ensure_ascii=False)
                outfile.write('\n')
            except json.JSONDecodeError:
                # 如果遇到无法解析的行，跳过
                continue

if __name__ == "__main__":
    input_file = './application.log'
    output_file = './formatted_application.log'
    format_log_file(input_file, output_file)
    print(f"Formatted log file written to {output_file}")
