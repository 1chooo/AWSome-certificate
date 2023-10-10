# 1027 Workshop - Run Down

## Overview

## Lambda Part
1. 我們先登錄進 AWS Console 並且在上方欄位輸入搜尋 lambda
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

