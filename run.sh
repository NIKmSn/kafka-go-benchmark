go build
records=( 10000 50000 100000 200000 500000 )
for i in "${records[@]}"
do
  echo
  echo "-------------------------------------PRODUCE[$i]-------------------------------------------"
  ./kafka-go-benchmark -driver=segmentio -strategy=produce -records=$i
  ./kafka-go-benchmark -driver=sarama -strategy=produce -records=$i
  echo "-------------------------------------CONSUME-----------------------------------------------"
  ./kafka-go-benchmark -driver=segmentio -strategy=consume -records=$i
  ./kafka-go-benchmark -driver=sarama -strategy=consume -records=$i
  echo "-------------------------------------FINISHED----------------------------------------------"
  echo
done