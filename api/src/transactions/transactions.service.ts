import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { TransactionServiceInterface } from './transaction.service.interface';
import { TransactionDto } from './transaction.dto';
import { ClientKafka } from '@nestjs/microservices';

@Injectable()
export class TransactionsService
  implements TransactionServiceInterface, OnModuleInit
{
  constructor(
    @Inject('KAFKA_SERVICE') private readonly kafkaClient: ClientKafka,
  ) {}
  async onModuleInit() {
    await this.kafkaClient.connect();
  }
  createTransaction(transactionDto: TransactionDto) {
    this.kafkaClient.emit('transactions', JSON.stringify(transactionDto));
  }
}
