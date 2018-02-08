package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func main() {
	//http: //blog.studygolang.com/2012/12/%E6%A0%87%E5%87%86%E5%BA%93-xml%E5%A4%84%E7%90%86%EF%BC%88%E4%B8%80%EF%BC%89/
	var t xml.Token
	var err error

	input := `<Order xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
	<Where></Where>
	<Mobile>18122101755</Mobile>
	<Phone>18122101755</Phone>
	<Email></Email>
	<Province>19</Province>
	<City>1601</City>
	<County>3633</County>
	<Name></Name>
	<Id>0</Id>
	<Pin></Pin>
	<TypeId>0</TypeId>
	<IdProvince>19</IdProvince>
	<IdCity>1601</IdCity>
	<IdArea>3633</IdArea>
	<IdTown>0</IdTown>
	<IsPromiseShipment>false</IsPromiseShipment>
	<SelectedBankId>0</SelectedBankId>
	<SelectedPeriod>0</SelectedPeriod>
	<CouponDiscount>0</CouponDiscount>
	<PromotionPrice>2949.00</PromotionPrice>
	<DiscountLipinka>0</DiscountLipinka>
	<TotalFee>0</TotalFee>
	<MoneyBalance>0</MoneyBalance>
	<IsUseBalance>false</IsUseBalance>
	<IsPopEmpty>true</IsPopEmpty>
	<OrderGuid>0f07db29-96ba-456b-a856-2afb7f23cb51</OrderGuid>
	<Price>2999.00</Price>
	<Discount>50.00</Discount>
	<OrderId>71925538574</OrderId>
	<ShipTime>0</ShipTime>
	<UserOrderId>71925554311</UserOrderId>
	<IdPaymentType>4</IdPaymentType>
	<IdShipmentType>70</IdShipmentType>
	<RfType>0</RfType>
	<IdInvoiceType>1</IdInvoiceType>
	<IdInvoiceHeaderType>4</IdInvoiceHeaderType>
	<IdInvoiceContentsType>1</IdInvoiceContentsType>
	<IdInvoicePutType>3</IdInvoicePutType>
	<IsPutBookInvoice>false</IsPutBookInvoice>
	<IdInvoiceContentTypeBook>4</IdInvoiceContentTypeBook>
	<CODTime>3</CODTime>
	<CodDate>2018-02-02T00:00:00</CodDate>
	<IsCodInform>false</IsCodInform>
	<PaymentWay>0</PaymentWay>
	<IdPickSite>0</IdPickSite>
	<IdPickSiteName></IdPickSiteName>
	<IsInv>true</IsInv>
	<InvoiceTitle>个人</InvoiceTitle>
	<TheLipinkas/>
	<IdSchool>0</IdSchool>
	<IdCompanyBranch>10</IdCompanyBranch>
	<IdDelivery>0</IdDelivery>
	<ServerName>host-10-191-175-31</ServerName>
	<OrderType>22</OrderType>
	<DiscountMobile>0</DiscountMobile>
	<DiscountTelecomIntegral>0</DiscountTelecomIntegral>
	<IsUseTelecomIntegral>false</IsUseTelecomIntegral>
	<UnionAlternationDays>0</UnionAlternationDays>
	<UnionId>0</UnionId>
	<UnionTime>0001-01-01T00:00:00</UnionTime>
	<ClientIP>113.65.13.212</ClientIP>
	<UserId>68215874</UserId>
	<IsRegister>false</IsRegister>
	<UserLevel>90</UserLevel>
	<userScore>201</userScore>
	<RePrice>0.00</RePrice>
	<IdBank>0</IdBank>
	<PayType>0</PayType>
	<BidId>0</BidId>
	<Ver>0</Ver>
	<IsNewOrder>false</IsNewOrder>
	<StoreId>110007548</StoreId>
	<ParentId>71925554311</ParentId>
	<FromId>0</FromId>
	<SplitType>2</SplitType>
	<PayMoney>2949.00</PayMoney>
	<Executor></Executor>
	<OrderBulk>47498.0170</OrderBulk>
	<IsJdShip>true</IsJdShip>
	<Tags>
	  <int>3011</int>
	  <int>2013</int>
	</Tags>
	<SendPay>00000000110000000000000002001000030100160000000000600000000001010000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200020000000010</SendPay>
	<TheExt/>
	<TheExts/>
	<TheExtTags>
	  <ExtTagPair>
		<Key>port</Key>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>SOP_VENDER_SENDPAY</Key>
		<Val>162705:36:1</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>aerfaTime</Key>
		<Val>{&quot;apt&quot;:[{&quot;venderId&quot;:162705,&quot;codDate&quot;:null,&quot;batch&quot;:null,&quot;codTime&quot;:&quot;3&quot;,&quot;carrier&quot;:null,&quot;idSopShipmentType&quot;:68,&quot;offset&quot;:0,&quot;batchId&quot;:0,&quot;menDianId&quot;:0,&quot;stationId&quot;:0},{&quot;venderId&quot;:532446,&quot;codDate&quot;:null,&quot;batch&quot;:null,&quot;codTime&quot;:&quot;3&quot;,&quot;carrier&quot;:&quot;[]&quot;,&quot;idSopShipmentType&quot;:67,&quot;offset&quot;:0,&quot;batchId&quot;:0,&quot;menDianId&quot;:0,&quot;stationId&quot;:0}],&quot;dpt&quot;:[]}</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>__jdb</Key>
		<Val>138814566.35.15163309934331207366780|35.1517553261</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>mt_xid</Key>
		<Val>V2_52007VwMWUlpaUFsXQRtZB2UDEVNbWlVdGUAYbAxlBkBWXVxXRh0ZGV0ZYgERV0FRUw8YVRBcBTRWFlUIX1EJH3kaXQVuHxNSQVtVSx9LEl4EbAYRYl9oUmoXSBxaAW4FF1JUXlVdHEEbWgBjMxFaXQ%3D%3D</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>newInvoiceVersion</Key>
		<Val>0</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>sdkServerName</Key>
		<Val>10.190.165.64</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>__jda</Key>
		<Val>138814566.15163309934331207366780.1516330993.1517551314.1517553261.35</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>DisChannels</Key>
		<Val>PC_LF</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>shopIdFlags</Key>
		<Val>[]</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>invoice_kaipiaofangshi</Key>
		<Val>0</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>plusUserFlag</Key>
		<Val>0000000000000000000000000000010900100100000000216300000000000000000000000000000000000000000000000000</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>addressDetail</Key>
		<Val></Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>addressId</Key>
		<Val>40404335</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>forceBotTag</Key>
		<Val>000</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>__jdv</Key>
		<Val></Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>add_lng</Key>
		<Val>0.0</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>add_tag</Key>
		<Val>0</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>promise311tag</Key>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>unpl</Key>
		<Val>%3d%3d</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>AutoCancelTime</Key>
		<Val>345600</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>PromoUuid</Key>
		<Val>PROMO-b09c4ae9-d84c-4e15-8d56-2f203cdbf7d8</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>mt_subsite</Key>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>add_lat</Key>
		<Val>0.0</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>outerIds</Key>
		<Val>{&quot;200100123903&quot;:&quot;^bf8e31c9bcbd468c901c0e8441c8c08b&quot;}</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>isDefaultAddress</Key>
		<Val>2</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>splitGiftRelation</Key>
		<Val>2</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>Chuna</Key>
		<Val>uniSystemN</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>mt_ext</Key>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>AutoShowCancelTime</Key>
		<Val>345600</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>PaySureDate</Key>
		<Val>2018-02-02T14:56:00</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>newSopShopTag</Key>
		<Val>2015</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>limitBuyUUID</Key>
		<Val>71925554311</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>AutoCancelTimeSource</Key>
		<Val>2</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>SendPayDict</Key>
		<Val>33:3|38:1|39:6|8:1|9:1|196:0|199:0|198:1|50:6|185:2|189:2|25:2|95:1|63:1|61:1</Val>
	  </ExtTagPair>
	  <ExtTagPair>
		<Key>Rpc_Remote_SOP_Freight</Key>
		<Val>-1</Val>
	  </ExtTagPair>
	</TheExtTags>
	<CSOrderType>0</CSOrderType>
	<InitFactPrice>2949.00</InitFactPrice>
	<Status>6</Status>
	<Status2>9</Status2>
	<Printx>0</Printx>
	<Ziti>0</Ziti>
	<Yn>1</Yn>
	<IsForAddCoupon>false</IsForAddCoupon>
	<IsSubsidy>false</IsSubsidy>
	<CreateDate>2018-02-02T14:54:52</CreateDate>
	<ScareBuyState>0</ScareBuyState>
	<Pop>
	  <VenderId>162705</VenderId>
	  <VenderName></VenderName>
	  <VenderType>0</VenderType>
	  <AreaNum>1</AreaNum>
	</Pop>
	<froms/>
	<IsUseScore>false</IsUseScore>
	<JiFenCount>0</JiFenCount>
	<JiFenDiscountMoney>0</JiFenDiscountMoney>
	<bigItemCodTime>0</bigItemCodTime>
	<PaymentInfoList>
	  <OrderPaymentInfo>
		<IdPaymentType>4</IdPaymentType>
		<SupportedSkuIdList>
		  <long>17744395513</long>
		</SupportedSkuIdList>
		<SupportedSkuUuidList>
		  <string>10_7109472689_162705_17744395513_1_0</string>
		</SupportedSkuUuidList>
	  </OrderPaymentInfo>
	</PaymentInfoList>
	<ShipmentInfoList>
	  <OrderShipmentInfo>
		<IdShipmentType>70</IdShipmentType>
		<SupportedSkuIdList>
		  <long>17744395513</long>
		</SupportedSkuIdList>
		<IsJdShip>true</IsJdShip>
		<PaymentWay>0</PaymentWay>
		<CodDate>2018-02-02T00:00:00</CodDate>
		<CodTime>3</CodTime>
		<IsCodInform>false</IsCodInform>
		<IdPickSite>0</IdPickSite>
		<SupportedSkuUuidList>
		  <string>10_7109472689_162705_17744395513_1_0</string>
		</SupportedSkuUuidList>
	  </OrderShipmentInfo>
	</ShipmentInfoList>
	<FreightInfo/>
	<DParentId>71925554311</DParentId>
	<IsShcoolPick>false</IsShcoolPick>
	<JingDouCount>0</JingDouCount>
	<JingDouRate>100</JingDouRate>
	<ZZweight>0</ZZweight>
	<SXZZweight>0</SXZZweight>
  </Order>8e5220aa_bbc2_4eb9_b849_a4afd76bc6e3<Cart xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
	<Num>21</Num>
	<HasDIY>false</HasDIY>
	<HasDaJiaDian>false</HasDaJiaDian>
	<HasBook>false</HasBook>
	<Tags>
	  <int>1</int>
	</Tags>
	<TotalPrice>2999.00</TotalPrice>
	<Weight>0.0</Weight>
	<TotalDiscount>50.00</TotalDiscount>
	<TotalRePrice>0.00</TotalRePrice>
	<CanUseDQ>true</CanUseDQ>
	<OrderId>71925554311</OrderId>
	<TheGifts>
	  <Suit>
		<MainSKU>
		  <Id>17744395513</Id>
		  <Cid>720</Cid>
		  <Name></Name>
		  <Weight>11.0000</Weight>
		  <ImgUrl>jfs/t10384/72/972747391/130015/57a0e7bd/59dae020Nfbeb370a.jpg</ImgUrl>
		  <cid2>716</cid2>
		  <Bilv>1</Bilv>
		  <YbChildId>0</YbChildId>
		  <YbChildNum>0</YbChildNum>
		  <IsYb>false</IsYb>
		  <YbEntitys/>
		  <IsScareBuy>false</IsScareBuy>
		  <Price>2999.00</Price>
		  <Discount>50.00</Discount>
		  <Score>20</Score>
		  <Jing>0</Jing>
		  <rePrice>0.00</rePrice>
		  <Num>1</Num>
		  <NumLimit>1</NumLimit>
		  <SkuBulk>47498.0170</SkuBulk>
		  <Maxpurchqty>0</Maxpurchqty>
		  <PromoId>7109472689</PromoId>
		  <Stock>5</Stock>
		  <AwardType>0</AwardType>
		  <IdDelivery>0</IdDelivery>
		  <ShopId>162705</ShopId>
		  <ShopName></ShopName>
		  <ShopType>0</ShopType>
		  <Tags/>
		  <TheExtTags>
			<ExtTagPair>
			  <Key>YBForCenter</Key>
			  <Val>[]</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>mobileStoreId</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>isJX</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>PromotionSkuTypes</Key>
			  <Val>[{&quot;qt&quot;:1,&quot;id&quot;:7081743134,&quot;pty&quot;:1,&quot;pdc&quot;:50.0,&quot;ptg&quot;:0,&quot;tid&quot;:&quot;0&quot;,&quot;mnl&quot;:0,&quot;dl&quot;:0.0,&quot;minn&quot;:1},{&quot;qt&quot;:0,&quot;id&quot;:7109472689,&quot;pty&quot;:4,&quot;pdc&quot;:0.0,&quot;ptg&quot;:0,&quot;tid&quot;:&quot;0&quot;,&quot;mnl&quot;:0,&quot;dl&quot;:0.0,&quot;minn&quot;:0}]</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>isFBPRevenue</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>skuMark</Key>
			  <Val>32</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>wserver</Key>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>sfkc</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>skuUuid</Key>
			  <Val>10_7109472689_162705_17744395513_1_0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>AgriGoods</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>BrandID</Key>
			  <Val>172761</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>StoreId</Key>
			  <Val>-1</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>EDIskuExtParam</Key>
			  <Val>00000000000000000000</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>isVMI</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>shangjia</Key>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>NewStoreId</Key>
			  <Val>110007548</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>originId</Key>
			  <Val>1</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>itemId</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>additionalWeightSku</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>v</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>isFreeInstall</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>skuModel</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>goodstype</Key>
			  <Val>0</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>nationallySetWare</Key>
			  <Val>3</Val>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>YStype</Key>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>JPLY</Key>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>outerId</Key>
			</ExtTagPair>
			<ExtTagPair>
			  <Key>suitNo</Key>
			  <Val>1</Val>
			</ExtTagPair>
		  </TheExtTags>
		  <PromoMemPriceList/>
		</MainSKU>
		<TheSkus/>
		<Id>7109472689</Id>
		<IsScareBuy>false</IsScareBuy>
		<Type>0</Type>
		<Price>2999.00</Price>
		<Discount>50.00</Discount>
		<Num>1</Num>
		<Score>0</Score>
		<DonatedScore>0.0</DonatedScore>
		<RequiredScore>0.0</RequiredScore>
		<RePrice>0.00</RePrice>
		<SuitType>0</SuitType>
		<NeedMoney>0</NeedMoney>
		<AddMoney>0</AddMoney>
		<AwardNum>0</AwardNum>
		<Money>0</Money>
		<Jing>0</Jing>
	  </Suit>
	</TheGifts>
	<PromoInfo>1@{ts:[1000],aid:&quot;19-1601-3633-0&quot;,ul:[{id:&quot;17744395513&quot;,n:1,tl:[],cl:[],r:{amount:0.00},p:{amount:2999.00},d:{amount:50.00},sl:[{pid:7081743134,pt:1,ps:0,mfmzFullRefundType:0},{pid:-100,pt:1,ps:0,mfmzFullRefundType:0},{pid:7109472689,pt:4,ps:0,mfmzFullRefundType:0},{pid:-200,pt:4,ps:0,mfmzFullRefundType:0}],s:20,pl:[{et:1518234540000,tk:&quot;0&quot;,pid:7081743134,pt:1,cl:[],r:{amount:0.00},d:{amount:50.00},s:20,lut:4,buyLimit:1,bt:1517278518000,pg:0,mnl:0,dl:{amount:0.00}},{et:1523064600000,tk:&quot;0&quot;,pid:7109472689,pt:4,cl:[],r:{amount:0.00},d:{amount:0.00},s:0,lut:0,buyLimit:1,bt:1517473647000,pg:0,mnl:0,dl:{amount:0.00}}],gl:[{t:2,id:&quot;200100123903&quot;,n:20,p:{amount:0.00},d:{amount:0.00},gs:0,pt:0}],gpl:[],jdPrice:&quot;2949.00&quot;}],sl:[],cl:[],csi:{td:{amount:50.00},tp:{amount:2999.00},tr:{amount:0.00},ts:20,cl:[],tn:21},ip:&quot;10.190.48.255&quot;}</PromoInfo>
  </Cart>`
	inputReader := strings.NewReader(input)

	// 从文件读取，如可以如下：
	// content, err := ioutil.ReadFile("studygolang.xml")
	// decoder := xml.NewDecoder(bytes.NewBuffer(content))

	decoder := xml.NewDecoder(inputReader)
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		// 处理元素结束（标签）
		case xml.EndElement:
			fmt.Printf("Token of '%s' end\n", token.Name.Local)
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
			// ...
		}
	}
}
