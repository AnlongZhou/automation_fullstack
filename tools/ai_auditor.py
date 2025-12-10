import sys
from openai import OpenAI

client = OpenAI(
    base_url='http://localhost:11434/v1',
    api_key='ollama',
)


def analyze_code(file_path):
    try:
        with open(file_path, 'r') as file:
            code_content = file.read()
    except FileNotFoundError:
        print(f"Error: couldn't find file {file_path}")
        sys.exit(1)

    print(f"Sending {file_path} to local Ollama Agent")

    try:
        response = client.chat.completions.create(
            model="llama3.2",
            messages=[
                {
                    "role": "system",
                    "content": "You are a Senior Golang Security Engineer. Analyze the code provided. "
                               "1. Provide a 1-sentence summary. "
                               "2. List any critical bugs or security issues. "
                               "3. If safe, say '✅ Code looks good.' Concise."
                },
                {
                    "role": "user",
                    "content": f"Code:\n\n{code_content}"
                }
            ]
        )

        print("\n")
        print("Ollama report")
        print(response.choices[0].message.content)
        print("\n")

    except Exception as e:
        print(f"Error communicating with Ollama: {e}")


if __name__ == "__main__":
    analyze_code("main.go")
