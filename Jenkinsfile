pipeline {
    agent any

    tools {
        maven 'Maven 3'
    }

    stages {

        stage('Клонирование проекта') {
            steps {
                git branch: 'lab4_v2', url: 'https://github.com/Stepan714/PDRIS_Sber.git'
            }
        }

        stage('Тестирование') {
            steps {
                sh '''
                pytest app/tests --cov=app --cov-report xml:coverage.xml --alluredir=allure-results || true
                '''
            }
        }

        stage('Проверка кодстайла') {
            steps {
                withSonarQubeEnv('SonarQube') {
                    sh 'mvn clean sonar:sonar'
                }
            }
        }

        stage('Баги в Allure') {
            steps {
                allure([
                    reportBuildPolicy: 'ALWAYS',
                    results: [[path: 'allure-results']]
                ])
            }
        }

        
        stage('Поднятие приложения') {
            steps {
                sh 'cd ansible && ansible-playbook -i inventory deploy.yml'
            }
        }
    }

    post {
        always {
            echo 'Пайплан полностью выполнен'
        }
    }
}
