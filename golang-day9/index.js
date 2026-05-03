// burda hoist meselesi var demeleli , 
// var i elan olunduğu üçün, yaddaşda cəmi bir dənə i qutusu olur. 
// Loop hər dəfə dönəndə həmin o tək qutunun içindəki rəqəmi dəyişir.
//  setTimeout daxilindəki funksiya "mən bir az sonra işləyəcəm" deyib kənara çəkilir.
// Loop çox sürətlə işləyib bitir. Loop bitəndə i-nin həmin o tək qutudakı son dəyəri 10 olur.
// Bir az sonra setTimeout funksiyaları işə düşəndə hamısı qaçıb həmin o tək qutuya baxır. 
// Qutuda isə artıq 10 rəqəmi var. Ona görə də hamısı ekrana 10 çıxarır


for (var i = 0; i < 10; i++) {
  setTimeout(function() {
    console.log(i);
  }, 100);
}
