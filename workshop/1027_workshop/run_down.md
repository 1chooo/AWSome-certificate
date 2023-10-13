# 1027 Workshop - Run Down

## Overview

## Lambda Part
1. 我們先登錄進 AWS Console 並且在上方欄位輸入搜尋 lambda [x] 
2. 進入 lambda 頁面後我們點擊按鈕以建立「lambda 函式」
3. 選擇「從頭開始撰寫」，輸入 function name 並且 runtime 選擇 python3.9，架構選擇 x86_64
4. 進入我們剛剛創立的 lambda 函數
5. 我們先來帶大家使用 lambda 函數成功執行的樣子，點擊 test 後會進入 Configure test event。
6. 我們先暫時將 EventName 命名為 TestEvent ，接著點擊 save，之後會回到程式碼的畫面，我們再次點擊 Test 按鈕，以查看執行果
7. 接著我們會到一個 Environment Variable 的頁籤，如果 Response 裡的 StatusCode 顯示 200 代表我們的 lambda 能夠成功執行，那通常我們在程式開發中只要 StatusCode 不是顯示 4 開頭的大部分都代表我們的程式可以正常運作
8. 接著我們準備貼上 jeopardy 遊戲的程式碼已進行我們遊戲玩法的設定。
9.  我們要先刪除剛剛的測試 Event ，我們需要點擊 Test 旁邊的小箭頭，會出現一個 Configure test event 的按鈕，大家請幫我點擊，進去頁面後，會出現剛剛的 TestEvent 名稱右邊有個刪除按鈕大家放心直接點擊並且選擇永久刪除，最後點擊 x 回到剛剛的 lambda_function.py 畫面。
10. 我們可以先回到 lambda_function.py 的頁籤將我們提供的程式碼貼上，透過 ctrl/command + A 的方式全選起來做貼上，以確保之前的程式碼完全覆蓋成 jeopardy 的遊戲邏輯。
11. 貼上 jeopardy 的遊戲邏輯後我們就要再定義一個新的 Event ，我們將之命名為 GeoEvent ，另外大家幫我貼上預先準備好的 json 內容，同樣將下方的 Event JSON 全選起來直接貼上，我們最後點擊 save。
12. 那 Hugo 這邊要提醒大家，我們在使用 lambda 函數時，非常重要的一個細節，Test 按鈕的旁邊還有一個按鈕，名稱為 Deploy ，我們只要在有修改任何 lambda function 的情況下，我們一定要點擊一次 deploy（deploy 按鈕旁也有一句標語提示我們 changes not deployed）， 否則將不會執行我們最新更新的程式碼，那這邊舉個小例子幫助大家記憶，如果我們今天在寫一份文件，我們想要讓電腦記得我們剛剛編輯的痕跡，我們就會做出存檔的動作，而同樣我們在編輯 lambda function 時，也要做出存檔的動作，以運行我們最新編寫好的程式。
13. 那我們點擊 deploy 後 lambda 將記得我們的最新程式，於是我們再點擊一次 Test 按鈕。
14. 這時候發現結果並非如我們的預期，出現 "Unable to import module 'lambda_function': No module named 'requests'" 的錯誤，大家不要害怕，有時候的出錯反而是讓我們通向成功的道路，於是我們閱讀一下這則錯誤訊息，原來是因為我們還缺少了一些步驟，那接下我將帶大家解決這個問題。
15. 那接下來大家要幫我下載 python.zip 的這個檔案，切記大家下載下來後我們保持原本的檔名 python.zip
16. 接著幫我點開左邊的側邊欄位，在最後有個 layers 的選項，需要大家幫我點擊去。
17. 進入後右上角有個 Create Layer 的選項，會進入下個頁面
18. 我們將名稱命名為 jeopardy_request ，下方有個上傳的按鈕，大家切記，上方要保持勾選 Upload a .zip file，上傳我們剛剛下載的 python.zip 檔案，並且下方幫我繼續勾選 x86_64 ， runtime 也要選擇 python3.9，最後我們點擊 create。
19. 完成 add a layer 後我們回到剛剛建立的 function ，幫我將此頁面下拉，最下面有個 layers 的欄位，大家幫我點擊 add a layer 
20. 大家幫我點選 custom layers ，選擇我們剛剛建立的 jeopardy_request layer 並且版本幫我選擇 1。
21. 最後也就是最後一步啦！大家直接點擊 Test 若出現以下畫面就大功告成啦！我們就可以開始設定我們 lex 機器人了！

## Lex Part

1. 在上方欄位輸入搜尋 Lex
2. 進到畫面後點選 Create bot，選擇 create a blank bot 名稱輸入 jeopardy_2023_1027_workshop，描述是註記給自己看的，因此大家可以填寫幫助自己記憶的內容，如果不太清楚要怎麼填寫可以填上，workshop demo on 2023/10/27，I am role 選擇 Create a role with basic Amazon Lex permissions. Children's Online Privacy Protection Act (COPPA) 選擇 No 因為我們並未涉及孩童相關的內容，填寫完以上後點擊 Next 。
3. 我們的語言選擇 English，Voice interaction 大家可以根據自己的喜好做選擇，有非常多種可以做嘗試，Intent classification confidence score threshold 保持 0.4
4. 接著我們下滑找到 Intent details 並且取名為 play_new_round，描述填寫 start the new round。
5. 我們到 Sample utterances ，我們新增三個描述 Let's play the Jeopardy!, Start a Jeopardy., Have fun with Jeopardy.
6. 再來我們要設定機器人的起始回應，我們往下找到 Initial response，輸入 Okay, let's have fun! 最後我們點擊 save 儲存。
7. 我們已經完成機器人的初始化了！接著我們要繼續新增 intent 以讓我們的對話更為豐富，，因此我們點擊側邊的 intent，我們再新增 intent ，因此要點擊 add intent ，並且命名為 check_score
8. 接著我們會進入到 check_score intent 的畫面，這是為了讓我們查詢遊戲的分數
9. 我們要到 Sample utterances 輸入 grade, my score, check score
10. 同樣我們要設定機器人的起始回應，我們往下找到 Initial response，輸入 checking score is not available
11. 最後我們點擊 save intent。
12. 這時候我們點擊 build 來測試我們剛剛設計的機器人對話是否成功吧！這時候上方繪有藍條的進度條
13. 當顯示成功 build 的進度條後我們就可以點擊 Test 來測試我們的 Bot
14. 我們可以測試剛剛設定的對話機器人是否有回應我們希望的內容，因此我們可以在側邊的聊天框輸入剛剛設定好的 Let's play the Jeopardy!, Start a Jeopardy., Have fun with Jeopardy, grade, my score, check score 看看是否出現 Okay, let's have fun!, checking score is not available ，另外我們也可以測試剛剛沒有設定的對話機器人會怎麼回覆
15. 若是我們沒有設定的對話會出現 Intent FallbackIntent is fulfilled ，這是 Lex 的預設回應，如果再傳送他不認識的訊息，則會這樣回覆我們。
16. 接下來我們會需要新增 slot 位置類型，點選 Add slot type 我們要選擇空白的 (add blank slot type)，並且命名為 Category。
17. 下一步我們需要把 Slot value resolution 設定為 Restrict to slot values，接著找到 Slot type values ，位置類型值我們需要新增三個項目，透過 add value 的方法新增三個種類，最後我們點擊儲存。
18. 在我們設定好 slot type 後，接下來我們要繼續設定我們的機器人，我們先回到一開始設定的 intent: play_new_round 我們現在要新增 slot。
19. 點擊 add a slot 後，我們需要輸入「名稱、類型以及提示」分別是 CategoryName, Category, Choose you question type，最後點擊儲存。
20. 接著我們往下滑找到 confirmation，我們要在 Confirmation prompt 填入 make sure your question type：{CategoryName} Decline response 填入 cancel question type：{CategoryName}
21. 接著我們需要找到 fulfillment 的 On successful fulfillment 填入 ok,preparing {CategoryName} question。最後我們點擊儲存 intent。
22. 現在我們已經設定好機器人的回覆機制了，我們現在要幫一開始設立的 lambda 遊戲邏輯加入我們的機器人，我們同樣停留在 play_new_round，我們往下尋找 add a fulfillment ，找到 advanced option，把 Use a Lambda function for fulfillment 的選項勾選，因為我們要加入我們剛剛設定好的 lambda function，點選 update options，以及 save intent。
23. 最後我們點選 build 來建立我們的機器人，如果成功後我們點擊 Test，我們進到對話框，點選右上角的 setting 按鈕，將 lambda function 的 source 調整成我們剛設立好的 lambda function ，並且點擊 save。
24. 最後我們就可以在對話框看到我們剛剛設定的結果啦！



