import pytest
from app import app

@pytest.fixture
def client():
    app.config['TESTING'] = True
    with app.test_client() as client:
        yield client

def test_increment(client):
    response = client.post('/', data={'increment': 2})
    assert b'Новое значение: 2' in response.data

