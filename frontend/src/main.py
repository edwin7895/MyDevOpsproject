from flask import Flask, render_template, request, redirect, url_for

app = Flask(__name__)

@app.route('/')
def home():
    return render_template('index.html')

@app.route('/about')
def about():
    return render_template('about.html')

@app.route('/contact', methods=['GET', 'POST'])
def contact():
    if request.method == 'POST':
        # Capturar los datos del formulario
        name = request.form['name']
        email = request.form['email']
        message = request.form['message']
        
        # Aquí puedes agregar lógica para procesar los datos, como enviarlos por correo
        # o almacenarlos en una base de datos. Por ahora, solo mostramos un mensaje de éxito.
        
        success_message = f"Thank you, {name}. Your message has been received."
        return render_template('contact.html', success=success_message)
    
    return render_template('contact.html')

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5001)
