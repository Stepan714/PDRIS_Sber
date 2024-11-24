pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Убедитесь, что здесь указан правильный URL вашего репозитория
                git 'https://github.com/Stepan714/PDRIS_Sber.git'
            }
        }

        stage('Build') {
            steps {
                sh 'pip install -r requirements.txt'
            }
        }

        stage('Test') {
            steps {
                sh 'pytest --junitxml=test-results.xml'
            }
        }

        stage('SonarQube analysis') {
            environment {
                scannerHome = tool 'SonarQube Scanner'
            }
            steps {
                withSonarQubeEnv('SonarQube') {
                    sh "${scannerHome}/bin/sonar-scanner"
                }
            }
        }

        stage('Check Quality Gate') {
            steps {
                script {
                    def qg = waitForQualityGate()
                    if (qg.status != 'OK') {
                        error "Pipeline aborted due to quality gate failure: ${qg.status}"
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                ansiblePlaybook credentialsId: 'your-ansible-credentials-id',
                                playbook: 'deploy.yml'
            }
        }
    }

    post {
        always {
            junit 'test-results.xml'
            archiveArtifacts artifacts: 'test-results.xml', allowEmptyArchive: true
        }
    }
}
