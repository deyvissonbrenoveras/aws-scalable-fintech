import { TransactionDto } from './transaction.dto';

export interface TransactionServiceInterface {
  createTransaction(transactionDto: TransactionDto): void;
}
