import { Component, OnInit} from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { UserService } from '../services/user/user.service';
import { UseraccountsService } from '../services/useraccounts/useraccounts.service';

@Component({
    selector: 'app-transaction',
    templateUrl: './transaction.component.html',
    styleUrls: ['./transaction.component.scss']
})
export class TransactionComponent implements OnInit {
    user: any = null;

    date = null;
    day = null;
    month = null;
    year = null

    beneficiary_account: number = null;
    amount: number = null;

    constructor(
        private route: ActivatedRoute,
        private userAccountsService: UseraccountsService,
    ) { }

    ngOnInit(): void {
        this.route.data
            .subscribe((data: {user: any}) => {
                this.user = data.user.data;
            });

        this.date = new Date();
        this.day = this.date.getDate();
        this.month = Number(this.date.getMonth()) + 1;
        this.year = this.date.getFullYear();
    }

    onKey(event: any, type: string) {
        if (type === 'beneficiary_account') {
            this.beneficiary_account = event.target.value;
        }

        if (type === 'amount') {
            this.amount = event.target.value;
        }
    }

    onTransaction(): void {
        this.userAccountsService
            .transaction(this.beneficiary_account, this.amount);


    }
}
