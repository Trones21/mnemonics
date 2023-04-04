import {Component} from '@angular/core'
import { IonicModule } from '@ionic/angular'


@Component({
    selector: 'app-single-collection',
    templateUrl: 'singleCollection.page.html',
    standalone:true,
    imports:[
        IonicModule
    ]
  })
export class SingleCollectionPage{
    collection = {
        name: "Math",
        logoPath:'',
        id: "6425fe604e175c2a4a8265e9",
        stars: 51,
        views: 216,
        createdAt: "2011-09-04T06:31:12.212Z",
        updatedAt: "2011-09-05T06:31:12.212Z",
        mnemonics:[
            {
                name:"PEMDAS",
                hint:"Please Excuse My Dear Aunt Sally",
                answer:"PEMDAS - Parentheses, Exponents, Multiplication/Division (Left to Right), Addition/Subtraction (Left to Right)",
                stars:2,
                views:500,
                createdAt: "2011-09-04T06:31:12.212Z",
                updatedAt: "2011-09-05T06:31:12.212Z"
            },
            {
                name:"FOIL",
                hint:"FOIL - Distributing binomials",
                answer:`In order to get the ax2+bx+c form, we must FOIL out (2x−5)2.
                FOIL is technique for distributing two binomials. The letters stand for First, Outer, Inner, and Last.
                
                "First" stands for multiply the terms which occur first in each binomial.
                
                "Outer" stands for multiply the outermost terms in the product.
                
                "Inner" stands for multiply the innermost two terms.
                
                "Last" stands for multiply the terms which occur last in each binomial."`,
                stars:10,
                views:100,
                createdAt: "2011-09-04T06:31:12.212Z",
                updatedAt: "2011-09-05T06:31:12.212Z"
            }
        ]

    }

    calculateCollectionRanks(){
        for(let mnemonic of this.collection.mnemonics){

        }
    }
}