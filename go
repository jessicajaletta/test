<?

set_time_limit(0);

while(true){

$rand = rand(1,10);
$pollid = "10293235";
$voting_id = "47357756";

$useragent="Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.1) Gecko/20061204 Firefox/2.0.0.1";

$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, "http://answers.poll.fm/vote/?va=10&pt=0&r=".$rand."&p=".$pollid."&a=".$voting_id);
curl_setopt($ch, CURLOPT_HEADER, true);
curl_setopt($ch, CURLOPT_USERAGENT, $useragent); 
$res = curl_exec ($ch) ;


curl_close($ch);

sleep(6);
}
?>
