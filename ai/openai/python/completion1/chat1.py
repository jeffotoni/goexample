from openai import OpenAI
client = OpenAI()

completion = client.chat.completions.create(
  model="gpt-3.5-turbo",
  messages=[
    {"role": "system", "content": "Qual a melhor linguagem de programação do mundo?"},
    #{"role": "user", "content": ""}
  ]
)

print(completion.choices[0].message)

