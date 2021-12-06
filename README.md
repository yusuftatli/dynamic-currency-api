# en ucuz kur bilgisini dönen api

belirlediğimiz süre aralıklarla job üzerinden tanımları yapılan providerlara eşzamanlı requestler atılıp channel vasıtası ile dönen değerler arasında en ucuz kur bilgisi hesaplanıp redis üzerine yazılmaktadır. burada provider sayısını bağımlılık durumu en aza indirilecek düzeyde tasarlanmıştır. yeni eklenecek providerlerin dönüş tipinin farklı olabilmesi, auth gerekliliği gibi durumlar dikkate alınmıştır. performans öncelikli olduğu için api isteklerini direk redis üzerinden kur bilgisi okunarak dönülmektedir. 

## api request 
http://localhost:8081/currency/EURO
result: {
    "Currency": "EURO",
    "Data": 11.4
}

http://localhost:8081/currency/USD
result: {
    "Currency": "USD",
    "Data": 9.96
}

http://localhost:8081/currency/GBP
result: {
    "Currency": "GBP",
    "Data": 13.31
}

```bash
pip install foobar
```

## Usage

```golang
make install
make run
```




## License
[MIT](https://choosealicense.com/licenses/mit/)