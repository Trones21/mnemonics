
//     function rankBy(objArr, propertytoRankOn, rankPropertyName){

//         //other idea: set rank to 0, pull items out of the objArr, and push them back in before returning with like null or something 
        
//         for(let [idx, item] of objArr.entries()){
//             if(!item.hasOwnProperty(propertytoRankOn)){
//                 console.error(`Object missing ${propertytoRankOn} property`)
//                 return false;
//             }
//         }
            
//         objArr = objArr.sort((a, b) => (a[propertytoRankOn] < b[propertytoRankOn]) ? 1 : -1)
        
            
//         let rank = 0;
        
//         for(let [idx,item] of objArr.entries()){
//             //always update
//             rank += 1;
        
//             if(idx === 0){
//                 objArr[idx][rankPropertyName] = 1
//                 continue;
//             }
//             if(item[propertytoRankOn] === sorted[idx - 1][propertytoRankOn]){
//                 objArr[idx][rankPropertyName]  = -1
//             }
        
//             if(item[propertytoRankOn] !== sorted[idx - 1][propertytoRankOn]){
//                 objArr[idx][rankPropertyName] = rank;
//             }
//         }
        
        
//         //Apply Ties 
//         for(let [idx,item] of objArr.entries()){
//             if(objArr[idx][rankPropertyName]  === -1){   
//                 let retval = getNextRank(idx, objArr)
//                 item[rankPropertyName]  = retval
//             }
//         }
        
        
//         return objArr
            
//         function getNextRank(idx, arr){
//             if(arr[idx][rankPropertyName]  === -1){
//                 //recurse / other oher returns
//                 return getNextRank(idx + 1, arr)
//             }
//             //innermost return
//             return arr[idx][rankPropertyName] 
//         }
        
//         }
// }
