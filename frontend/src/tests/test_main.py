import unittest
from app import app

class FlaskAppTests(unittest.TestCase):

    # Se ejecuta antes de cada prueba
    def setUp(self):
        # Crea un cliente de pruebas
        self.app = app.test_client()
        self.app.testing = True

    # Prueba para la ruta de inicio
    def test_home_page(self):
        response = self.app.get('/')
        self.assertEqual(response.status_code, 200)
        self.assertIn(b'Home', response.data)

    # Prueba para la ruta 'about'
    def test_about_page(self):
        response = self.app.get('/about')
        self.assertEqual(response.status_code, 200)
        self.assertIn(b'About', response.data)

    # Prueba para la ruta 'contact' con GET
    def test_contact_page_get(self):
        response = self.app.get('/contact')
        self.assertEqual(response.status_code, 200)
        self.assertIn(b'Contact', response.data)

    # Prueba para la ruta 'contact' con POST
    def test_contact_page_post(self):
        response = self.app.post('/contact', data=dict(
            name='Test User',
            email='test@example.com',
            message='This is a test message.'
        ))
        self.assertEqual(response.status_code, 200)
        self.assertIn(b'Thank you, Test User', response.data)

if __name__ == '__main__':
    unittest.main()
