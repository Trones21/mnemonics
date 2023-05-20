import { Injectable } from '@angular/core';
import { DataService } from '../data.service';

@Injectable({
  providedIn: 'root'
})
export class MnemonicService {
  baseUrl: string
  constructor(dataService: DataService,) { 
     this.baseUrl = dataService.baseUrl;
  }

  async getMnemonic(mnemonicId: string){
    let res = await fetch(`${this.baseUrl}/mnemonic/${mnemonicId}`)
    let obj = await res.json();
    console.log(obj.profile)
    return obj.profile;
  }

  async getMnemonics(mnemonicIds: Array<string>){
    let idsStr = ''
    for(let id of mnemonicIds){
      idsStr += id + ','
    }
    idsStr = idsStr.slice(0, idsStr.lastIndexOf(','))
    let res = await fetch(`${this.baseUrl}/mnemonic?ids=[${idsStr}]`)
    let obj = await res.json();
    console.log(obj.profile)
    return obj.profile;
  }

  async createMnemonic(mnemonic: any){
    let res = await fetch(`${this.baseUrl}/mnemonic?new`,{
        method:'POST',
        body: JSON.stringify(mnemonic),
        headers:{
            'Content-Type':'application/json'
        }
    })
    let obj = await res.json();
    console.log(obj.item)
    return 'Created Object';
  }

}
