pipeline {
    // 使用 Golang 官方 Docker 映像作為構建和測試環境，確保環境隔離和一致性
    agent {
        docker {
            image 'golang:1.22'
            // 允許容器內產生文件 (例如 allure-results) 且權限正確
            args '-u 0:0'
        }
    }

    // 環境變數：設定 Allure 報告的輸出路徑
    environment {
        ALLURE_RESULTS_PATH = 'allure-results'
    }

    stages {
        // ----------------------------------------
        // 1. 環境準備 (在 Docker 容器內)
        // ----------------------------------------
        stage('Prepare Environment') {
            steps {
                echo 'Cleaning up existing allure results directory...'
                // 確保報告目錄乾淨
                sh "rm -rf ${env.ALLURE_RESULTS_PATH}"
                sh "mkdir -p ${env.ALLURE_RESULTS_PATH}"
            }
        }

        // ----------------------------------------
        // 2. 依賴與構建
        // ----------------------------------------
        stage('Install Dependencies & Build') {
            steps {
                echo 'Installing Go dependencies...'
                // 下載並安裝專案依賴，特別是 allure-go 庫
                sh 'go mod tidy'

                // 執行 Go 專案構建（如果您的整合測試需要執行檔）
                // 這裡我們假設只需要測試檔案，但保留 Build 階段是好習慣
                // sh 'go build -o my_app .'
            }
        }

        // ----------------------------------------
        // 3. 執行整合測試
        // ----------------------------------------
        stage('Run Integration Tests') {
            steps {
                echo 'Running Allure-Go Integration Tests...'
                // 關鍵步驟：設定 ALLURE_RESULTS_PATH 環境變數 (已在 environment 區塊設定)
                // 執行 go test，它會自動將 JSON 報告寫入 ALLURE_RESULTS_PATH
                sh "go test -v ./..."
            }
        }

        // ----------------------------------------
        // 4. 生成並發布 Allure Report
        // ----------------------------------------
        stage('Publish Allure Report') {
            steps {
                echo 'Generating and Publishing Allure Report...'

                // 使用 Allure Jenkins Plugin 提供的指令
                // results: 必須指向存放 JSON 檔案的目錄 (即 ALLURE_RESULTS_PATH)
                allure([
                    includeProperties: false,
                    jdk: '',
                    properties: [],
                    reportBuildPolicy: 'ALWAYS',
                    results: [[path: env.ALLURE_RESULTS_PATH]]
                ])

                // 為了方便除錯，如果需要，可以將原始 JSON 文件作為 Artifacts 歸檔
                archiveArtifacts artifacts: "${env.ALLURE_RESULTS_PATH}/*.json", onlyIfSuccessful: false
            }
        }
    }

    // ----------------------------------------
    // 報告後處理
    // ----------------------------------------
    post {
        always {
            // 清理工作區，釋放 Jenkins Agent 上的空間
            cleanWs()
        }
        success {
            echo 'Pipeline finished successfully. Check the Allure Report link!'
        }
        failure {
            echo 'Pipeline failed. Check test logs.'
        }
    }
}