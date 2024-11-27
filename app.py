from flask import Flask, request
from redis import Redis

app = Flask(__name__)
redis = Redis(host='redis', port=6379)

@app.route('/', methods=['GET', 'POST'])
def index():
    if request.method == 'POST':
        increment = request.form.get('increment', 1, type=int)
    else:
        increment = 1

    previous_value = int(redis.get('hits') or 0)
    new_value = previous_value + increment
    redis.set('hits', new_value)

    content = f"""
    <html>
      <head>
        <style>
          .center {{
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            text-align: center;
            font-family: Arial, sans-serif;
          }}
          .box {{
            border: 2px solid black;
            padding: 20px;
            margin: 10px;
          }}
          form {{
            margin-top: 20px;
          }}
          h1 {{
            margin-bottom: 20px;
          }}
          input[type='number'] {{
            width: 50px;
          }}
        </style>
      </head>
      <body>
        <div class="center">
          <div>
            <h1>Счетчик посещений</h1>
            <div class="box">Старое значение: <strong>{previous_value}</strong></div>
            <div class="box">Новое значение: <strong>{new_value}</strong></div>
            <form method="POST">
              <label for="increment">Увеличить на:</label>
              <input type="number" id="increment" name="increment" value="1" min="1">
              <button type="submit">Обновить</button>
            </form>
          </div>
        </div>
      </body>
    </html>
    """
    return content

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80)
