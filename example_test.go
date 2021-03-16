package xmldom_test

import (
	"fmt"
	"github.com/kehula/go-xmldom"
	"log"
)

const (
	ExampleXml = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE junit SYSTEM "junit-result.dtd">
<testsuites>
	<testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/go-xmldom">
		<properties>
			<property name="go.version">go1.8.1</property>
		</properties>
		<testcase classname="go-xmldom" id="ExampleParseXML" time="0.004"></testcase>
		<testcase classname="go-xmldom" id="ExampleParse" time="0.005"></testcase>
	</testsuite>
</testsuites>`
)

var xmlText = "<ns2:ImportNotificationsOfOrderExecutionRequest xmlns=\"urn:dom.gosuslugi.ru/common/1.0.0\" test=\"123\" xmlns:ns2=\"urn:dom.gosuslugi.ru/payments/1.0.0\" xmlns:ns3=\"urn:dom.gosuslugi.ru/payments-common/1.0.0\">\n\t\t\t\t\t\t<information-system-id>2c891fbb-a5bf-45ef-8c89-f82c7060c9d3</information-system-id>\n\t\t\t\t\t\t<organization-id>729faabe-d5ef-4eb8-a45c-e5aea1a9276e</organization-id>\n\t\t\t\t\t\t<ns2:notification-of-order-execution>\n\t\t\t\t\t\t\t<transport-id>14a5442a-f81c-4b04-82ca-b91d16691770</transport-id>\n\t\t\t\t\t\t\t<ns2:recipient-info>\n\t\t\t\t\t\t\t\t<ns2:inn>5247004695</ns2:inn>\n\t\t\t\t\t\t\t\t<ns2:legal>\n\t\t\t\t\t\t\t\t\t<ns2:kpp>000000000</ns2:kpp>\n\t\t\t\t\t\t\t\t\t<ns2:name>Открытое акционерное общество   Выксунский металлургический завод</ns2:name>\n\t\t\t\t\t\t\t\t</ns2:legal>\n\t\t\t\t\t\t\t\t<ns2:payment-information>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-inn>5247004695</ns3:recipient-inn>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-kpp>000000000</ns3:recipient-kpp>\n\t\t\t\t\t\t\t\t\t<ns3:bank-name>ПАО АКБ МЕТАЛЛИНВЕСТБАНК</ns3:bank-name>\n\t\t\t\t\t\t\t\t\t<ns3:payment-recipient>Открытое акционерное общество   Выксунский металлургический завод</ns3:payment-recipient>\n\t\t\t\t\t\t\t\t\t<ns3:bik>044525176</ns3:bik>\n\t\t\t\t\t\t\t\t\t<ns3:operating-account-number>40702810500000000470</ns3:operating-account-number>\n\t\t\t\t\t\t\t\t</ns2:payment-information>\n\t\t\t\t\t\t\t</ns2:recipient-info>\n\t\t\t\t\t\t\t<ns2:order-info>\n\t\t\t\t\t\t\t\t<ns2:order-id>40445251762502210000000000000034</ns2:order-id>\n\t\t\t\t\t\t\t\t<ns2:order-date>2021-02-25+03:00</ns2:order-date>\n\t\t\t\t\t\t\t\t<ns2:amount>2493151</ns2:amount>\n\t\t\t\t\t\t\t\t<ns2:payment-purpose>Уплата процентов по договору банковского счета 98/810-Х Доп.1020 от 24/02/2021 Без НДС</ns2:payment-purpose>\n\t\t\t\t\t\t\t\t<ns2:payment-document-number>000000000000000000000000000668</ns2:payment-document-number>\n\t\t\t\t\t\t\t</ns2:order-info>\n\t\t\t\t\t\t</ns2:notification-of-order-execution>\n\t\t\t\t\t\t<ns2:notification-of-order-execution>\n\t\t\t\t\t\t\t<transport-id>910dc7e1-e618-46f1-9301-0dae2ae3047c</transport-id>\n\t\t\t\t\t\t\t<ns2:recipient-info>\n\t\t\t\t\t\t\t\t<ns2:inn>3662178961</ns2:inn>\n\t\t\t\t\t\t\t\t<ns2:legal>\n\t\t\t\t\t\t\t\t\t<ns2:kpp>366201001</ns2:kpp>\n\t\t\t\t\t\t\t\t\t<ns2:name>ООО ВоронежТеплоГазСервис</ns2:name>\n\t\t\t\t\t\t\t\t</ns2:legal>\n\t\t\t\t\t\t\t\t<ns2:payment-information>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-inn>3662178961</ns3:recipient-inn>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-kpp>366201001</ns3:recipient-kpp>\n\t\t\t\t\t\t\t\t\t<ns3:bank-name>ЦЕНТРАЛЬНО-ЧЕРНОЗЕМНЫЙ БАНК ПАО СБЕРБАНК</ns3:bank-name>\n\t\t\t\t\t\t\t\t\t<ns3:payment-recipient>ООО ВоронежТеплоГазСервис</ns3:payment-recipient>\n\t\t\t\t\t\t\t\t\t<ns3:bik>042007681</ns3:bik>\n\t\t\t\t\t\t\t\t\t<ns3:operating-account-number>40702810513000063063</ns3:operating-account-number>\n\t\t\t\t\t\t\t\t</ns2:payment-information>\n\t\t\t\t\t\t\t</ns2:recipient-info>\n\t\t\t\t\t\t\t<ns2:order-info>\n\t\t\t\t\t\t\t\t<ns2:order-id>40445251762502210000000000000033</ns2:order-id>\n\t\t\t\t\t\t\t\t<ns2:order-date>2021-02-25+03:00</ns2:order-date>\n\t\t\t\t\t\t\t\t<ns2:amount>800000</ns2:amount>\n\t\t\t\t\t\t\t\t<ns2:payment-purpose>Оплата за обслуживание за февраль 2021по счету 00бп-000307 от 01.02,21 Без НДС</ns2:payment-purpose>\n\t\t\t\t\t\t\t\t<ns2:payment-document-number>000000000000000000000000000668</ns2:payment-document-number>\n\t\t\t\t\t\t\t</ns2:order-info>\n\t\t\t\t\t\t</ns2:notification-of-order-execution>\n\t\t\t\t\t\t<ns2:notification-of-order-execution>\n\t\t\t\t\t\t\t<transport-id>358f04e0-f8d8-48c9-9f82-c9aa0ea228c4</transport-id>\n\t\t\t\t\t\t\t<ns2:recipient-info>\n\t\t\t\t\t\t\t\t<ns2:inn>3123110760</ns2:inn>\n\t\t\t\t\t\t\t\t<ns2:legal>\n\t\t\t\t\t\t\t\t\t<ns2:kpp>312301001</ns2:kpp>\n\t\t\t\t\t\t\t\t\t<ns2:name>ОАО Белгородская сбытовая компания</ns2:name>\n\t\t\t\t\t\t\t\t</ns2:legal>\n\t\t\t\t\t\t\t\t<ns2:payment-information>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-inn>3123110760</ns3:recipient-inn>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-kpp>312301001</ns3:recipient-kpp>\n\t\t\t\t\t\t\t\t\t<ns3:bank-name>БЕЛГОРОДСКОЕ ОТДЕЛЕНИЕ N8592 ПАО СБЕРБАНК</ns3:bank-name>\n\t\t\t\t\t\t\t\t\t<ns3:payment-recipient>ОАО Белгородская сбытовая компания</ns3:payment-recipient>\n\t\t\t\t\t\t\t\t\t<ns3:bik>041403633</ns3:bik>\n\t\t\t\t\t\t\t\t\t<ns3:operating-account-number>40702810607070102063</ns3:operating-account-number>\n\t\t\t\t\t\t\t\t</ns2:payment-information>\n\t\t\t\t\t\t\t</ns2:recipient-info>\n\t\t\t\t\t\t\t<ns2:order-info>\n\t\t\t\t\t\t\t\t<ns2:order-id>40445251762502210000000000000032</ns2:order-id>\n\t\t\t\t\t\t\t\t<ns2:order-date>2021-02-25+03:00</ns2:order-date>\n\t\t\t\t\t\t\t\t<ns2:amount>1000000</ns2:amount>\n\t\t\t\t\t\t\t\t<ns2:payment-purpose>Оплата за свет по счету б/н от 22.02.21  В т.ч. НДС 20% - 1666.67</ns2:payment-purpose>\n\t\t\t\t\t\t\t\t<ns2:payment-document-number>000000000000000000000000000668</ns2:payment-document-number>\n\t\t\t\t\t\t\t</ns2:order-info>\n\t\t\t\t\t\t</ns2:notification-of-order-execution>\n\t\t\t\t\t\t<ns2:notification-of-order-execution>\n\t\t\t\t\t\t\t<transport-id>2d05a18f-7880-43fd-bb0a-aedb5bb0ac72</transport-id>\n\t\t\t\t\t\t\t<ns2:recipient-info>\n\t\t\t\t\t\t\t\t<ns2:inn>5247004695</ns2:inn>\n\t\t\t\t\t\t\t\t<ns2:legal>\n\t\t\t\t\t\t\t\t\t<ns2:kpp>000000000</ns2:kpp>\n\t\t\t\t\t\t\t\t\t<ns2:name>Открытое акционерное общество   Выксунский металлургический завод</ns2:name>\n\t\t\t\t\t\t\t\t</ns2:legal>\n\t\t\t\t\t\t\t\t<ns2:payment-information>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-inn>5247004695</ns3:recipient-inn>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-kpp>000000000</ns3:recipient-kpp>\n\t\t\t\t\t\t\t\t\t<ns3:bank-name>ПАО АКБ МЕТАЛЛИНВЕСТБАНК</ns3:bank-name>\n\t\t\t\t\t\t\t\t\t<ns3:payment-recipient>Открытое акционерное общество   Выксунский металлургический завод</ns3:payment-recipient>\n\t\t\t\t\t\t\t\t\t<ns3:bik>044525176</ns3:bik>\n\t\t\t\t\t\t\t\t\t<ns3:operating-account-number>40702810500000000470</ns3:operating-account-number>\n\t\t\t\t\t\t\t\t</ns2:payment-information>\n\t\t\t\t\t\t\t</ns2:recipient-info>\n\t\t\t\t\t\t\t<ns2:order-info>\n\t\t\t\t\t\t\t\t<ns2:order-id>40445251762502210000000000000031</ns2:order-id>\n\t\t\t\t\t\t\t\t<ns2:order-date>2021-02-25+03:00</ns2:order-date>\n\t\t\t\t\t\t\t\t<ns2:amount>56958904</ns2:amount>\n\t\t\t\t\t\t\t\t<ns2:payment-purpose>Уплата процентов по депозиту по Соглашению ДДЮ-077-1 от 05/07/2013 и Подтверждению 3838       от 24/02/2021. Без НДС.</ns2:payment-purpose>\n\t\t\t\t\t\t\t\t<ns2:payment-document-number>000000000000000000000000000668</ns2:payment-document-number>\n\t\t\t\t\t\t\t</ns2:order-info>\n\t\t\t\t\t\t</ns2:notification-of-order-execution>\n\t\t\t\t\t\t<ns2:notification-of-order-execution>\n\t\t\t\t\t\t\t<transport-id>b3483636-6856-4c2e-9e0f-21854497a5c4</transport-id>\n\t\t\t\t\t\t\t<ns2:recipient-info>\n\t\t\t\t\t\t\t\t<ns2:inn>5247004695</ns2:inn>\n\t\t\t\t\t\t\t\t<ns2:legal>\n\t\t\t\t\t\t\t\t\t<ns2:kpp>000000000</ns2:kpp>\n\t\t\t\t\t\t\t\t\t<ns2:name>Открытое акционерное общество   Выксунский металлургический завод</ns2:name>\n\t\t\t\t\t\t\t\t</ns2:legal>\n\t\t\t\t\t\t\t\t<ns2:payment-information>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-inn>5247004695</ns3:recipient-inn>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-kpp>000000000</ns3:recipient-kpp>\n\t\t\t\t\t\t\t\t\t<ns3:bank-name>ПАО АКБ МЕТАЛЛИНВЕСТБАНК</ns3:bank-name>\n\t\t\t\t\t\t\t\t\t<ns3:payment-recipient>Открытое акционерное общество   Выксунский металлургический завод</ns3:payment-recipient>\n\t\t\t\t\t\t\t\t\t<ns3:bik>044525176</ns3:bik>\n\t\t\t\t\t\t\t\t\t<ns3:operating-account-number>40702810500000000470</ns3:operating-account-number>\n\t\t\t\t\t\t\t\t</ns2:payment-information>\n\t\t\t\t\t\t\t</ns2:recipient-info>\n\t\t\t\t\t\t\t<ns2:order-info>\n\t\t\t\t\t\t\t\t<ns2:order-id>40445251762502210000000000000030</ns2:order-id>\n\t\t\t\t\t\t\t\t<ns2:order-date>2021-02-25+03:00</ns2:order-date>\n\t\t\t\t\t\t\t\t<ns2:amount>540000000000</ns2:amount>\n\t\t\t\t\t\t\t\t<ns2:payment-purpose>Возврат депозита по Соглашению ДДЮ-077-1 от 05/07/2013 и Подтверждению 3838       от 24/02/2021. Без НДС.</ns2:payment-purpose>\n\t\t\t\t\t\t\t\t<ns2:payment-document-number>000000000000000000000000000668</ns2:payment-document-number>\n\t\t\t\t\t\t\t</ns2:order-info>\n\t\t\t\t\t\t</ns2:notification-of-order-execution>\n\t\t\t\t\t\t<ns2:notification-of-order-execution>\n\t\t\t\t\t\t\t<transport-id>3171e6c7-2ce0-454c-880c-c65d33f5f9cd</transport-id>\n\t\t\t\t\t\t\t<ns2:recipient-info>\n\t\t\t\t\t\t\t\t<ns2:inn>1644040195</ns2:inn>\n\t\t\t\t\t\t\t\t<ns2:legal>\n\t\t\t\t\t\t\t\t\t<ns2:kpp>000000000</ns2:kpp>\n\t\t\t\t\t\t\t\t\t<ns2:name>ООО ТАТНЕФТЬ-АЗС ЦЕНТР</ns2:name>\n\t\t\t\t\t\t\t\t</ns2:legal>\n\t\t\t\t\t\t\t\t<ns2:payment-information>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-inn>1644040195</ns3:recipient-inn>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-kpp>000000000</ns3:recipient-kpp>\n\t\t\t\t\t\t\t\t\t<ns3:bank-name>ПАО БАНК ЗЕНИТ</ns3:bank-name>\n\t\t\t\t\t\t\t\t\t<ns3:payment-recipient>ООО ТАТНЕФТЬ-АЗС ЦЕНТР</ns3:payment-recipient>\n\t\t\t\t\t\t\t\t\t<ns3:bik>044525272</ns3:bik>\n\t\t\t\t\t\t\t\t\t<ns3:operating-account-number>40702810500730001540</ns3:operating-account-number>\n\t\t\t\t\t\t\t\t</ns2:payment-information>\n\t\t\t\t\t\t\t</ns2:recipient-info>\n\t\t\t\t\t\t\t<ns2:order-info>\n\t\t\t\t\t\t\t\t<ns2:order-id>40445251762502210000000000000029</ns2:order-id>\n\t\t\t\t\t\t\t\t<ns2:order-date>2021-02-25+03:00</ns2:order-date>\n\t\t\t\t\t\t\t\t<ns2:amount>1846000000</ns2:amount>\n\t\t\t\t\t\t\t\t<ns2:payment-purpose>Оплата за судовое топливо RMD по договору 20500/2020/1037/жд от 19.11.2020, согласно счета 478 от 19.02.21 Сумма 18460000-00 В т.ч. НДС  (20%) 3076666-67</ns2:payment-purpose>\n\t\t\t\t\t\t\t\t<ns2:payment-document-number>000000000000000000000000000668</ns2:payment-document-number>\n\t\t\t\t\t\t\t</ns2:order-info>\n\t\t\t\t\t\t</ns2:notification-of-order-execution>\n\t\t\t\t\t\t<ns2:notification-of-order-execution>\n\t\t\t\t\t\t\t<transport-id>19a34f3c-a5e9-48a2-8bca-8f170b82fc63</transport-id>\n\t\t\t\t\t\t\t<ns2:recipient-info>\n\t\t\t\t\t\t\t\t<ns2:inn>1644040195</ns2:inn>\n\t\t\t\t\t\t\t\t<ns2:legal>\n\t\t\t\t\t\t\t\t\t<ns2:kpp>000000000</ns2:kpp>\n\t\t\t\t\t\t\t\t\t<ns2:name>ООО ТАТНЕФТЬ-АЗС ЦЕНТР</ns2:name>\n\t\t\t\t\t\t\t\t</ns2:legal>\n\t\t\t\t\t\t\t\t<ns2:payment-information>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-inn>1644040195</ns3:recipient-inn>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-kpp>000000000</ns3:recipient-kpp>\n\t\t\t\t\t\t\t\t\t<ns3:bank-name>ПАО БАНК ЗЕНИТ</ns3:bank-name>\n\t\t\t\t\t\t\t\t\t<ns3:payment-recipient>ООО ТАТНЕФТЬ-АЗС ЦЕНТР</ns3:payment-recipient>\n\t\t\t\t\t\t\t\t\t<ns3:bik>044525272</ns3:bik>\n\t\t\t\t\t\t\t\t\t<ns3:operating-account-number>40702810500730001540</ns3:operating-account-number>\n\t\t\t\t\t\t\t\t</ns2:payment-information>\n\t\t\t\t\t\t\t</ns2:recipient-info>\n\t\t\t\t\t\t\t<ns2:order-info>\n\t\t\t\t\t\t\t\t<ns2:order-id>40445251762502210000000000000028</ns2:order-id>\n\t\t\t\t\t\t\t\t<ns2:order-date>2021-02-25+03:00</ns2:order-date>\n\t\t\t\t\t\t\t\t<ns2:amount>133276000</ns2:amount>\n\t\t\t\t\t\t\t\t<ns2:payment-purpose>Оплата за транспортные услуги по перевозке грузов по договору 20500/2020/1037/жд от 19.11.2020, согласно счета 480 от 19.02.21 Сумма 1332760-00 В т.ч. НДС  (20%) 222126-67</ns2:payment-purpose>\n\t\t\t\t\t\t\t\t<ns2:payment-document-number>000000000000000000000000000668</ns2:payment-document-number>\n\t\t\t\t\t\t\t</ns2:order-info>\n\t\t\t\t\t\t</ns2:notification-of-order-execution>\n\t\t\t\t\t\t<ns2:notification-of-order-execution>\n\t\t\t\t\t\t\t<transport-id>985e4705-26f5-491d-b9c4-ca87acdd1245</transport-id>\n\t\t\t\t\t\t\t<ns2:recipient-info>\n\t\t\t\t\t\t\t\t<ns2:inn>1644040195</ns2:inn>\n\t\t\t\t\t\t\t\t<ns2:legal>\n\t\t\t\t\t\t\t\t\t<ns2:kpp>000000000</ns2:kpp>\n\t\t\t\t\t\t\t\t\t<ns2:name>ООО ТАТНЕФТЬ-АЗС ЦЕНТР</ns2:name>\n\t\t\t\t\t\t\t\t</ns2:legal>\n\t\t\t\t\t\t\t\t<ns2:payment-information>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-inn>1644040195</ns3:recipient-inn>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-kpp>000000000</ns3:recipient-kpp>\n\t\t\t\t\t\t\t\t\t<ns3:bank-name>ПАО БАНК ЗЕНИТ</ns3:bank-name>\n\t\t\t\t\t\t\t\t\t<ns3:payment-recipient>ООО ТАТНЕФТЬ-АЗС ЦЕНТР</ns3:payment-recipient>\n\t\t\t\t\t\t\t\t\t<ns3:bik>044525272</ns3:bik>\n\t\t\t\t\t\t\t\t\t<ns3:operating-account-number>40702810500730001540</ns3:operating-account-number>\n\t\t\t\t\t\t\t\t</ns2:payment-information>\n\t\t\t\t\t\t\t</ns2:recipient-info>\n\t\t\t\t\t\t\t<ns2:order-info>\n\t\t\t\t\t\t\t\t<ns2:order-id>40445251762502210000000000000027</ns2:order-id>\n\t\t\t\t\t\t\t\t<ns2:order-date>2021-02-25+03:00</ns2:order-date>\n\t\t\t\t\t\t\t\t<ns2:amount>624000</ns2:amount>\n\t\t\t\t\t\t\t\t<ns2:payment-purpose>Оплата вознаграждения за организацию перевозок груза по договору 20500/2020/1037/жд от 19.11.2020, согласно счета 479 от 19.02.21 Сумма 6240-00 В т.ч. НДС  (20%) 1040-00</ns2:payment-purpose>\n\t\t\t\t\t\t\t\t<ns2:payment-document-number>000000000000000000000000000668</ns2:payment-document-number>\n\t\t\t\t\t\t\t</ns2:order-info>\n\t\t\t\t\t\t</ns2:notification-of-order-execution>\n\t\t\t\t\t\t<ns2:notification-of-order-execution>\n\t\t\t\t\t\t\t<transport-id>964f1e83-f4cf-4259-9d44-e96e9dc148cd</transport-id>\n\t\t\t\t\t\t\t<ns2:recipient-info>\n\t\t\t\t\t\t\t\t<ns2:inn>5247004695</ns2:inn>\n\t\t\t\t\t\t\t\t<ns2:legal>\n\t\t\t\t\t\t\t\t\t<ns2:kpp>000000000</ns2:kpp>\n\t\t\t\t\t\t\t\t\t<ns2:name>Акционерное общество   Выксунский металлургический завод</ns2:name>\n\t\t\t\t\t\t\t\t</ns2:legal>\n\t\t\t\t\t\t\t\t<ns2:payment-information>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-inn>5247004695</ns3:recipient-inn>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-kpp>000000000</ns3:recipient-kpp>\n\t\t\t\t\t\t\t\t\t<ns3:bank-name>ПАО АКБ МЕТАЛЛИНВЕСТБАНК</ns3:bank-name>\n\t\t\t\t\t\t\t\t\t<ns3:payment-recipient>Акционерное общество   Выксунский металлургический завод</ns3:payment-recipient>\n\t\t\t\t\t\t\t\t\t<ns3:bik>044525176</ns3:bik>\n\t\t\t\t\t\t\t\t\t<ns3:operating-account-number>40702810500000000470</ns3:operating-account-number>\n\t\t\t\t\t\t\t\t</ns2:payment-information>\n\t\t\t\t\t\t\t</ns2:recipient-info>\n\t\t\t\t\t\t\t<ns2:order-info>\n\t\t\t\t\t\t\t\t<ns2:order-id>40445251762502210000000000000026</ns2:order-id>\n\t\t\t\t\t\t\t\t<ns2:order-date>2021-02-25+03:00</ns2:order-date>\n\t\t\t\t\t\t\t\t<ns2:amount>14100</ns2:amount>\n\t\t\t\t\t\t\t\t<ns2:payment-purpose>Штраф за утерю временного пропуска без ндс</ns2:payment-purpose>\n\t\t\t\t\t\t\t\t<ns2:payment-document-number>000000000000000000000000000668</ns2:payment-document-number>\n\t\t\t\t\t\t\t</ns2:order-info>\n\t\t\t\t\t\t</ns2:notification-of-order-execution>\n\t\t\t\t\t\t<ns2:notification-of-order-execution>\n\t\t\t\t\t\t\t<transport-id>02ce7ec6-fbd2-4ec4-bf55-6441985ce4ae</transport-id>\n\t\t\t\t\t\t\t<ns2:recipient-info>\n\t\t\t\t\t\t\t\t<ns2:inn>3302001983</ns2:inn>\n\t\t\t\t\t\t\t\t<ns2:legal>\n\t\t\t\t\t\t\t\t\t<ns2:kpp>332801001</ns2:kpp>\n\t\t\t\t\t\t\t\t\t<ns2:name>МУП Владимирводоканал</ns2:name>\n\t\t\t\t\t\t\t\t</ns2:legal>\n\t\t\t\t\t\t\t\t<ns2:payment-information>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-inn>3302001983</ns3:recipient-inn>\n\t\t\t\t\t\t\t\t\t<ns3:recipient-kpp>332801001</ns3:recipient-kpp>\n\t\t\t\t\t\t\t\t\t<ns3:bank-name>ВЛАДИМИРСКОЕ ОТДЕЛЕНИЕ №8611 ПАО СБЕРБАНК</ns3:bank-name>\n\t\t\t\t\t\t\t\t\t<ns3:payment-recipient>МУП Владимирводоканал</ns3:payment-recipient>\n\t\t\t\t\t\t\t\t\t<ns3:bik>041708602</ns3:bik>\n\t\t\t\t\t\t\t\t\t<ns3:operating-account-number>40702810810040101515</ns3:operating-account-number>\n\t\t\t\t\t\t\t\t</ns2:payment-information>\n\t\t\t\t\t\t\t</ns2:recipient-info>\n\t\t\t\t\t\t\t<ns2:order-info>\n\t\t\t\t\t\t\t\t<ns2:order-id>40445251762502210000000000000025</ns2:order-id>\n\t\t\t\t\t\t\t\t<ns2:order-date>2021-02-25+03:00</ns2:order-date>\n\t\t\t\t\t\t\t\t<ns2:amount>108610</ns2:amount>\n\t\t\t\t\t\t\t\t<ns2:payment-purpose>Оплата за согласование топосъемки (дог. № 5782 от 09.02.21 г.) по сч. № 5782 от 09.02.21 г. В т.ч. НДС 20% - 181.02</ns2:payment-purpose>\n\t\t\t\t\t\t\t\t<ns2:payment-document-number>000000000000000000000000000668</ns2:payment-document-number>\n\t\t\t\t\t\t\t</ns2:order-info>\n\t\t\t\t\t\t</ns2:notification-of-order-execution>\n\t\t\t\t\t</ns2:ImportNotificationsOfOrderExecutionRequest>"

func ExampleParseNamespaces() {
	doc, err := xmldom.ParseXML(xmlText)
	if err != nil {
		log.Fatal(err)
	}
	for i, _ := range doc.NamespaceList {
		i++
	}
	fmt.Print("test")
	// Output:
	// test
}

func ExampleParseXML() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	fmt.Printf("name = %v\n", node.Name)
	fmt.Printf("attributes.len = %v\n", len(node.Attributes))
	fmt.Printf("children.len = %v\n", len(node.Children))
	fmt.Printf("root = %v", node == node.Root())
	// Output:
	// name = testsuites
	// attributes.len = 0
	// children.len = 1
	// root = true
}

func ExampleNode_GetAttribute() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	attr := node.FirstChild().GetAttribute("name")
	fmt.Printf("%v = %v\n", attr.Name, attr.Value)
	// Output:
	// name = github.com/subchen/go-xmldom
}

func ExampleNode_GetChildren() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	children := node.FirstChild().GetChildren("testcase")
	for _, c := range children {
		fmt.Printf("%v: id = %v\n", c.Name, c.GetAttributeValue("id"))
	}
	// Output:
	// testcase: id = ExampleParseXML
	// testcase: id = ExampleParse
}

func ExampleNode_FindByID() {
	root := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	node := root.FindByID("ExampleParseXML")
	fmt.Println(node.XML())
	// Output:
	// <testcase classname="go-xmldom" id="ExampleParseXML" time="0.004" />
}

func ExampleNode_FindOneByName() {
	root := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	node := root.FindOneByName("property")
	fmt.Println(node.XML())
	// Output:
	// <property name="go.version">go1.8.1</property>
}

func ExampleNode_FindByName() {
	root := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	nodes := root.FindByName("testcase")
	for _, node := range nodes {
		fmt.Println(node.XML())
	}
	// Output:
	// <testcase classname="go-xmldom" id="ExampleParseXML" time="0.004" />
	// <testcase classname="go-xmldom" id="ExampleParse" time="0.005" />
}

func ExampleNode_Query() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	// xpath expr: https://github.com/antchfx/xpath

	// find all children
	fmt.Printf("children = %v\n", len(node.Query("//*")))

	// find node matched tag name
	nodeList := node.Query("//testcase")
	for _, c := range nodeList {
		fmt.Printf("%v: id = %v\n", c.Name, c.GetAttributeValue("id"))
	}
	// Output:
	// children = 5
	// testcase: id = ExampleParseXML
	// testcase: id = ExampleParse
}

func ExampleNode_QueryOne() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	// xpath expr: https://github.com/antchfx/xpath

	// find node matched attr name
	c := node.QueryOne("//testcase[@id='ExampleParseXML']")
	fmt.Printf("%v: id = %v\n", c.Name, c.GetAttributeValue("id"))
	// Output:
	// testcase: id = ExampleParseXML
}

func ExampleDocument_XML() {
	doc := xmldom.Must(xmldom.ParseXML(ExampleXml))
	fmt.Println(doc.XML())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?><!DOCTYPE junit SYSTEM "junit-result.dtd"><testsuites><testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/go-xmldom"><properties><property name="go.version">go1.8.1</property></properties><testcase classname="go-xmldom" id="ExampleParseXML" time="0.004" /><testcase classname="go-xmldom" id="ExampleParse" time="0.005" /></testsuite></testsuites>
}

func ExampleDocument_XMLPretty() {
	doc := xmldom.Must(xmldom.ParseXML(ExampleXml))
	fmt.Println(doc.XMLPretty())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <!DOCTYPE junit SYSTEM "junit-result.dtd">
	// <testsuites>
	//   <testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/go-xmldom">
	//     <properties>
	//       <property name="go.version">go1.8.1</property>
	//     </properties>
	//     <testcase classname="go-xmldom" id="ExampleParseXML" time="0.004" />
	//     <testcase classname="go-xmldom" id="ExampleParse" time="0.005" />
	//   </testsuite>
	// </testsuites>
}

func ExampleNewDocument() {
	doc := xmldom.NewDocument("testsuites")

	testsuiteNode := doc.Root.CreateNode("testsuite").SetAttributeValue("name", "github.com/subchen/go-xmldom")
	testsuiteNode.CreateNode("testcase").SetAttributeValue("name", "case 1").Text = "PASS"
	testsuiteNode.CreateNode("testcase").SetAttributeValue("name", "case 2").Text = "FAIL"

	fmt.Println(doc.XMLPretty())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <testsuites>
	//   <testsuite name="github.com/subchen/go-xmldom">
	//     <testcase name="case 1">PASS</testcase>
	//     <testcase name="case 2">FAIL</testcase>
	//   </testsuite>
	// </testsuites>
}
