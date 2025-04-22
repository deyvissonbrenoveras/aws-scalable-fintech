import { Body, Controller, Inject, Post } from '@nestjs/common';
import { TransactionDto } from './transaction.dto';
import { TransactionsService } from './transactions.service';
import { TransactionServiceInterface } from './transaction.service.interface';

@Controller('transactions')
export class TransactionsController {
  constructor(
    @Inject(TransactionsService)
    private readonly transactionsService: TransactionServiceInterface,
  ) {}

  @Post()
  async createTransaction(@Body() transactionDto: TransactionDto) {
    await this.transactionsService.createTransaction(transactionDto);
  }
}
