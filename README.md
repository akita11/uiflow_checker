# uiflow_checker

WiFi環境が、[M5Stack](https://m5stack.com)の[UIFlow](https://flow.m5stack.com)で使用できるかをチェックします。

# 使い方

1. PCを、UIFlowで使う予定のWiFi（M5Stackの接続先）に接続します
2. PCで、"uiflow_checker"をダブルクリックで起動します。（Windows用：uiflow_checker_Win.exe、Mac(Intel)用：uiflow_checker_Mac、Mac(M1)用:uiflow_checker_MacM1)
3. チェックが行われます。エラーが出ていなければ、そのWiFiでUIFlowを使えます。何らかのエラーが表示されている場合は、そのWiFiではUIFlowを使えません。ネットワーク管理者に、flow.m5stack.comとの間でMQTT(TCP/1883)を通すよう、依頼してください。

# 技術詳細
UIFlowでは、デバイス(M5Stack)とサーバとの通信にMQTT(TCP/1883)を使います。WiFiの設定でこの通信を止めている場合は、UIFlow(Internet Mode)を利用できません。それを事前にチェックするツールです。

# Author

Junichi Akita (@akita11, akita@ifdl.jp)
