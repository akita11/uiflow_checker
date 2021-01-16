# uiflow_checker

[M5Stack](https://m5stack.com)の[UIFlow](https://flow.m5stack.com)で、デバイス側がサーバとの通信を行えるかどうかをチェックします。

# 使い方

1. PCを、UIFlowで使う予定のWiFiに接続します
2. PCで、"uiflow_checker"をダブルクリックで起動します。（Windows用：uiflow_checker_Win.exe、Mac(Intel)用：uiflow_checker_Mac、Mac(M1)用:uiflow_checker_MacM1)
3. チェックが行われます。エラーが出ていなければ、そのWiFiでUIFlowを使えます。何らかのエラーが表示されている場合は、そのWiFiではUIFlowを使えません。ネットワーク管理者に、flow.m5stack.comとの間でMQTT(TCP/1883)を通すよう、依頼してください。

# 技術詳細
UIFlowでは、デバイスとサーバとの通信にMQTT(TCP/1883)を使います。WiFiの設定で、この通信を止めている場合も多く、UIFlowを使えない場合もあります。

# Author

Junichi Akita (@akita11, akita@ifdl.jp)
