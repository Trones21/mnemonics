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
