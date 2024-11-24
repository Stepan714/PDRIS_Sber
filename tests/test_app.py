import pytest
import requests
from app import app

@pytest.fixture
def client():
    with app.test_client() as client:
        yield client

def test_index_get(client):
    response = client.get('/')
    assert response.status_code == 200
    assert b"Счетчик посещений" in response.data

def test_index_post(client):
    response = client.post('/', data={'increment': 5})
    assert response.status_code == 200
    assert b"Новое значение" in response.data

def test_default_increment(client):
    client.post('/')
    response = client.get('/')
    assert b"Новое значение: <strong>1</strong>" in response.data

def test_custom_increment(client):
    client.post('/', data={'increment': 10})
    response = client.get('/')
    assert b"Новое значение: <strong>11</strong>" in response.data

def test_multiple_increments(client):
    client.post('/', data={'increment': 3})
    client.post('/', data={'increment': 4})
    response = client.get('/')
    assert b"Новое значение: <strong>7</strong>" in response.data

