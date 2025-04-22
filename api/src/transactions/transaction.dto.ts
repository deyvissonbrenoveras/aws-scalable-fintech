import {
  IsEnum,
  IsInt,
  IsOptional,
  IsPositive,
  IsString,
  IsUUID,
} from 'class-validator';
import {
  TransactionMethodEnum,
  TransactionTypeEnum,
} from './transaction-enums.enum';

export class TransactionDto {
  @IsUUID('4')
  transactionId: string;

  @IsString()
  userId: string; // refactor to use userId from auth module

  @IsPositive()
  @IsInt()
  amount: number;

  @IsString()
  @IsOptional()
  description: string;

  @IsUUID('4')
  accountId: string;

  @IsEnum(TransactionTypeEnum)
  type: string;

  @IsEnum(TransactionMethodEnum)
  method: string;

  @IsString()
  destination: string;
}
