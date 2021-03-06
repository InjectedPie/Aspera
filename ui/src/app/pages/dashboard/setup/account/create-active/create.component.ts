import { Component, OnInit, OnDestroy, ViewChild, ViewEncapsulation } from '@angular/core';
import { Router } from '@angular/router';
import { MatStepper } from '@angular/material';
import { CreateService } from '../create.service';

@Component({
    selector: 'app-account-create',
    styleUrls: ['./create.component.scss'],
    templateUrl: './create.component.html'
})
export class CreateActiveAccountComponent implements OnInit {
    private stepIndex: number;
    @ViewChild(MatStepper) stepper: MatStepper;

    constructor(
        private router: Router,
        private createService: CreateService
    ) {}

    public ngOnInit() {

    }

    public ngOnDestroy() {

    }

    public setIndex() {
        this.stepper.selectedIndex = 1
    }

}
