
pipeline {
    agent any

    stages {
        stage('Клонирование проекта') {
            steps {
                git branch: 'lab4_v2', url: 'https://github.com/Stepan714/PDRIS_Sber.git'
            }
        }

        stage('Проверка содержимого') {
            steps {
                sh 'ls -la'
            }
        }

        stage('Тестирование') {
            steps {
                sh "pytest app/tests --cov=app --cov-report xml:coverage.xml --alluredir=allure-results || true"
            }
        }

        stage('Проверка кодстайла') {
            steps {
                withSonarQubeEnv('sonar') {
                    tool name: 'Maven', type: 'maven'
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
                dir('ansible') {
                    withEnv(["PATH+ANSIBLE=/usr/local/bin:/usr/bin"]) {
                        sh 'ansible-playbook -i inventory deploy.yml'
                    }
                }
            }
        }
    }

    post {
        always {
            echo 'Пайплайн полностью выполнен'
        }
    }
}
