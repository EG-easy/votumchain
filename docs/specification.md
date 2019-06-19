## Votum chain


### 背景
- 選挙や株式総会などの議決などの大多数による意思決定に際してのコスト削減。
- 投票結果の改ざんを防ぐ。
- 投票結果の照合が可能。
- community governanceの意思決定のあり方を提示する。
- 全国1741の市町村がそれぞれvalidatorを持つことが最終的にはよき。


### concept
重要事項を決定する際、およびその結果を公の場に公開する必要がある時に、誰でも使うことができるブロックチェーンアプリケーション。


### spec
- tendermint consensus
- voterは秘密鍵の管理？
- blockchain exploer


### stake holder
- validator
ブロックチェーンネットワークの維持

将来的には、地方自治体がvalidatorを立てることで公共的な性質をもつ。
- propser
投票の主催
- voter
投票を行う


### native token 
votum token
- 投票用のトークン発行時に必要
- 

### 手順
0. 投票内容・期間を決める
1. 投票用のトークンを発行(mint/distribute)
2. voterによる投票
3. 結果開示


voterは、addressを保有しておくことが必要。

### インセンティブ設計
透明性がある選挙制度を維持するためにvalidatorは市選管および国が持つべきか

そもそもそうするとvotum tokenは必要ない説濃厚

配布→無記名投票→シャッフルして匿名性の担保



## 今回のアプリケーションでの実装の順番

[x] 機能を決める

[ ] 状態(Keeper)を決める

[ ] Msgの中身を決める

[ ] cliを実装する

[ ] Rest apiを作る


## votum token model
proposal transactionを発行する(governance)
一定額をdepositする(deposit)
ここでproposerは投票用のトークンを発行(issure)
投票用トークンを配布する (distribute)
- 全てのアカウントに付与するor 特定のアカウントに付与する
- 配布する量を決めることができる
- votum tokenと同じ割合など
期限内で投票(vote)

すなわち、
- 既存のgovernance moduleを別のtokenでも使えるように改変する
- issure token, distribute tokenという別のmoduleを作る
-  一回delegateしないと投票できない仕組みになっている部分をどうするか



###


