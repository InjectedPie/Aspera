import { Injectable } from '@angular/core';
import { LoginPageActions } from '../../../../../auth/actions';
import { Store } from '@ngrx/store';
import * as fromAuth from '../../../../../auth/reducers';

@Injectable()
export class CreateService {
    private passphrase: string[];
    private pin: string;
    private id: string;
    private address: string;
    private stepIndex: number;

    constructor(private store: Store<fromAuth.State>) {
        this.stepIndex = 0;
        this.reset();
    }

    public setPassphrase(passphrase: string[]) {
        this.passphrase = passphrase;
    }

    public getPassphrase(): string[] {
        return this.passphrase;
    }

    public setPin(pin: string) {
        this.pin = pin;
    }

    public getPin(): string {
        return this.pin;
    }

    public getPassphrasePart(index: number): string {
        return this.passphrase[index];
    }

    public getCompletePassphrase(): string {
        return this.passphrase.join(" ");
    }

    public setId(id: string) {
        this.id = id;
    }

    public getId(id: string) {
        return this.id;
    }

    public setAddress(address: string) {
        this.address = address;
    }

    public getAddress() : string {
        return this.address;
    }

    public setStepIndex(index: number) {
        this.stepIndex = index;
    }

    public getStepIndex() : number {
        return this.stepIndex;
    }

    public isPassphraseGenerated() : boolean {
        return this.passphrase.length > 0 && this.address != undefined && this.id != undefined
    }

    public createAccount() {
        return this.store.dispatch(new LoginPageActions.Login({ credentials: { passphrase: this.getCompletePassphrase(), pin: this.pin } }));
    }

    public reset() {
        this.passphrase = [];
        this.id = undefined;
        this.address = undefined;
    }
}
