import {Component, Input} from '@angular/core'
import { IonicModule } from '@ionic/angular'
import { CollectionService } from './collection.service'
import { Mnemonic, MnemonicCard } from '../components/mnemonicCard'
import { CommonModule } from '@angular/common'
import { ActivatedRoute } from '@angular/router'

@Component({
    selector: 'app-single-collection',
    templateUrl: 'singleCollection.page.html',
    standalone: true,
    imports: [
        IonicModule,
        CommonModule,
        MnemonicCard
    ]
})
export class SingleCollectionPage{
    collection: Collection = {id: '1', name: 'Loading...'}
    constructor(private collectionService: CollectionService, private route: ActivatedRoute) {
    }

    ngOnInit(){
        let id: string = String(this.route.snapshot.paramMap.get('id'))
        if(id.length == 0){
            console.error('Unable to get id parameter from route')
        }
        this.collectionService.getCollectionByID(id, true).then(c => {
            
            if(!c.mnemonics){
                console.error('Collection mising mnemonics property (check spelling)', c)
            }
  
            this.collection = c;
            this.collection.mnemonicCount = this.collection.mnemonics?.length
            console.log(this.collection)

            //this.collection.mnemonicCards = (c.mnemonics as MnemonicCard)            
    })
}

    navBack(){
        console.warn("To implement -- need the id of the category")
    }
}



export interface Collection {
    id: string,
    name: string,
    stars?: number,
    views?: number,
    mnemonicCount?: number,
    mnemonics?: Mnemonic[]
}


function rankBy(objArr: any, propertytoRankOn: any, rankPropertyName: any){

    //other idea: set rank to 0, pull items out of the objArr, and push them back in before returning with like null or something 
    
    for(let [idx, item] of objArr.entries()){
        if(!item.hasOwnProperty(propertytoRankOn)){
            console.error(`Object missing ${propertytoRankOn} property`)
            return false;
        }
    }
        
    objArr = objArr.sort((a: any, b: any) => (a[propertytoRankOn] < b[propertytoRankOn]) ? 1 : -1)
    
        
    let rank = 0;
    
    for(let [idx,item] of objArr.entries()){
        //always update
        rank += 1;
    
        if(idx === 0){
            objArr[idx][rankPropertyName] = 1
            continue;
        }
        if(item[propertytoRankOn] === objArr[idx - 1][propertytoRankOn]){
            objArr[idx][rankPropertyName]  = -1
        }
    
        if(item[propertytoRankOn] !== objArr[idx - 1][propertytoRankOn]){
            objArr[idx][rankPropertyName] = rank;
        }
    }
    
    
    //Apply Ties 
    for(let [idx,item] of objArr.entries()){
        if(objArr[idx][rankPropertyName]  === -1){   
            let retval = getNextRank(idx, objArr)
            item[rankPropertyName]  = retval
        }
    }
    
    
    return objArr
        
    function getNextRank(idx: any, arr: Array<any>): Number{
        if(arr[idx][rankPropertyName]  === -1){
            //recurse / other oher returns
            return getNextRank(idx + 1, arr)
        }
        //innermost return
        return arr[idx][rankPropertyName] 
    }
    
}




            // c.mnemonics?.forEach( m => {
            //     let card: MnemonicCard = {

            //     }
        //             link: '232',
        //             name: m.name,
        //             stars: m.stars,
        //             views: m.views,
        //             globalStarRank: 15,
        //             globalViewRank: 17

        //         }
        //         let card = (mnemonic as MnemonicCard)
        //         this.collection.mnemonicCards = []
        //         this.collection.mnemonicCards = rankBy(this.collection.mnemonicCards, 'stars', 'collectionStarRank')
        //         this.collection.mnemonicCards = rankBy(this.collection.mnemonicCards, 'views', 'collectionViewRank')
        //         this.collection.mnemonicCards.push(card)
        //     })


        //     this.collection.mnemonics = (this.collection.mnemonics as Array<MnemonicCard>)
        //     console.log(this.collection)
        // })  
            
        
    //     // this.collection.mnemonics = rankBy(this.collection.mnemonics, 'stars', 'starRank')
    //     // this.collection.mnemonics = rankBy(this.collection.mnemonics, 'views', 'viewRank')
    // }
