<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>支付测试</title>
</head>
<script type="text/javascript">
  alert("sasdasdjlalsd");
</script>
<body>
<form id="alipaysubmit" name="alipaysubmit" action="https://mapi.alipay.com/gateway.do?_input_charset=utf-8" method="get" style='display:none;'> <input type="hidden" name="_input_charset" value="utf-8"> <input type="hidden" name="body" value="为你大爷充值99.80元"> <input type="hidden" name="notify_url" value="http://192.168.199.120:8088/payment?action=notufy"> <input type="hidden" name="out_trade_no" value="123456"> <input type="hidden" name="partner" value="2088411430532533"> <input type="hidden" name="payment_type" value="1"> <input type="hidden" name="return_url" value="http://192.168.199.120:8088/payment?action=return"> <input type="hidden" name="seller_email" value="759620527@qq.com"> <input type="hidden" name="service" value="create_direct_pay_by_user"> <input type="hidden" name="subject" value="充值100"> <input type="hidden" name="total_fee" value="99.80"> <input type="hidden" name="sign" value="KXml69tNZaQjE5tPcExPhOurUGrbPf5TQJanZ5WPbsLOKzMaz1sRsV%2BWtc53fRSUA
 %2FOvaWXJnwoSS%2BYJwJowtr3wGtwAMg4jNq7mPOOibYg63ef7%2B6VKC9emAyoL3Lc
 %2FHaYlGqAwPYCljjm%2BjL0zSSEHFZkfuMtvnFK1ZLNP238%3D"> <input type="hidden" name="sign_type" value="MD5"> </form> <script> document.forms['alipaysubmit'].submit(); </script>
</body>
</html>