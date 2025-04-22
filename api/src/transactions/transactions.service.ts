import { Injectable } from '@nestjs/common';
import { TransactionServiceInterface } from './transaction.service.interface';
import { TransactionDto } from './transaction.dto';

@Injectable()
export class TransactionsService implements TransactionServiceInterface {
  createTransaction(transactionDto: TransactionDto): Promise<void> {
    throw new Error('Method not implemented.');
  }
}
