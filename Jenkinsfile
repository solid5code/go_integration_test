pipeline {
    // 設置 Agent，這裡使用 Docker 代理，因為它更容易設置 Golang 環境
    agent {
        // 使用官方 Go 映像作為構建和測試環境
        docker {
            image 'golang:1.22'
            args '-u 0:0' // 確保文件權限問題被避免
        }
    }

    // 設置工具，確保 Jenkins 知道 go-junit-report 的路徑
    tools {
        // 假設 go-junit-report 已經被安裝到環境中
        // 如果您使用上面的 Docker agent，可以通過 sh 腳本安裝
    }

    stages {
        // ----------------------------------------
        // 1. 環境準備
        // ----------------------------------------
        stage('Environment Setup') {
            steps {
                sh 'go install github.com/jstemmer/go-junit-report/v2@latest'
            }
        }

        // ----------------------------------------
        // 2. 構建 (Build)
        // ----------------------------------------
        stage('Build') {
            steps {
                echo 'Starting Go project build...'
                // 根據您的專案需要執行構建，例如：
                sh 'go build -o my_app ./cmd/main.go'
                // 如果是微服務，可能還需要構建 Docker Image
            }
        }

        // ----------------------------------------
        // 3. 執行整合測試 (Integration Test)
        // ----------------------------------------
        stage('Run Integration Tests') {
            steps {
                echo 'Running Integration Tests and converting to JUnit-XML...'
                // 運行測試並使用 go-junit-report 轉換輸出
                // 這裡的 go-junit-report 命令會將結果導出到 allure-results/junit.xml
                // Allure 插件可以直接處理 JUnit-XML 檔案
                sh '''
                    mkdir -p allure-results
                    go test -v ./... | go-junit-report > allure-results/junit.xml
                '''
                // 如果您已經將命令包裝在腳本中，則執行：
                // sh './run_integration_tests.sh'
            }
        }

        // ----------------------------------------
        // 4. 生成 Allure Report (Report Generation)
        // ----------------------------------------
        stage('Generate Allure Report') {
            steps {
                // Archive the generated test results
                archiveArtifacts artifacts: 'allure-results/*.xml', onlyIfSuccessful: true

                // The Allure Jenkins Plugin generates the report
                // **注意:** results: 'allure-results' 必須指向包含測試結果文件 (.xml, .json, .properties, etc.) 的目錄
                allure([
                    includeProperties: false,
                    jdk: '',
                    properties: [],
                    reportBuildPolicy: 'ALWAYS',
                    results: [[path: 'allure-results']]
                ])
                echo 'Allure Report generated and published!'
            }
        }
    }

    // ----------------------------------------
    // 報告後處理 (Post-build Actions)
    // ----------------------------------------
    post {
        always {
            // 清理工作區
            cleanWs()
        }
        success {
            echo 'Pipeline finished successfully!'
        }
        failure {
            echo 'Pipeline failed. Check the logs for details.'
        }
    }
}