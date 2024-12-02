/* eslint-disable */
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Time: { input: any; output: any; }
  UUID: { input: any; output: any; }
};

export type DocUploadInput = {
  email: Scalars['String']['input'];
  url: Scalars['String']['input'];
  walletAddress: Scalars['String']['input'];
};

export type Land = {
  __typename?: 'Land';
  created_at: Scalars['Time']['output'];
  id: Scalars['UUID']['output'];
  size: Scalars['Int']['output'];
  symbol: Scalars['String']['output'];
  title: Scalars['String']['output'];
  titleDocument: Scalars['String']['output'];
  town?: Maybe<Scalars['String']['output']>;
  updated_at: Scalars['Time']['output'];
  verified: Verification;
};

export type Mutation = {
  __typename?: 'Mutation';
  chargeMpesa?: Maybe<Scalars['String']['output']>;
  updateTitleVerification: Title;
  uploadLandTitle: Title;
  uploadSupportingDoc: SupportingDoc;
};


export type MutationChargeMpesaArgs = {
  input: PayWithMpesaInput;
};


export type MutationUpdateTitleVerificationArgs = {
  input: UpdateTitleVerificationInput;
};


export type MutationUploadLandTitleArgs = {
  input: DocUploadInput;
};


export type MutationUploadSupportingDocArgs = {
  input: DocUploadInput;
};

export enum PaidFor {
  NotPaid = 'NOT_PAID',
  Paid = 'PAID'
}

export type PayWithMpesaInput = {
  currency: Scalars['String']['input'];
  email: Scalars['String']['input'];
  paymentFor: Scalars['UUID']['input'];
  phone: Scalars['String']['input'];
  reason: PaymentReason;
  walletAddress: Scalars['String']['input'];
};

export type Payment = {
  __typename?: 'Payment';
  created_at: Scalars['Time']['output'];
  id: Scalars['UUID']['output'];
  reference_id: Scalars['String']['output'];
  status: Scalars['String']['output'];
  title?: Maybe<Title>;
  title_id: Scalars['UUID']['output'];
  updated_at: Scalars['Time']['output'];
  wallet_address: Scalars['String']['output'];
};

export enum PaymentReason {
  LandRegistry = 'LAND_REGISTRY'
}

export enum PaymentStatus {
  Failed = 'failed',
  Pending = 'pending',
  Success = 'success',
  Timeout = 'timeout'
}

export type PaymentUpdate = {
  __typename?: 'PaymentUpdate';
  referenceId: Scalars['String']['output'];
  status: Scalars['String']['output'];
  walletAddress: Scalars['String']['output'];
};

export type Query = {
  __typename?: 'Query';
  getLocalLands: Array<Land>;
  getPaymentDetailsById: Payment;
  getPaymentsByStatus: Array<Payment>;
  getUserLands: Array<Land>;
};


export type QueryGetPaymentDetailsByIdArgs = {
  id: Scalars['UUID']['input'];
};


export type QueryGetPaymentsByStatusArgs = {
  status: PaymentStatus;
};


export type QueryGetUserLandsArgs = {
  walletAddress: Scalars['String']['input'];
};

export type Subscription = {
  __typename?: 'Subscription';
  paymentUpdate: PaymentUpdate;
};


export type SubscriptionPaymentUpdateArgs = {
  walletAddress: Scalars['String']['input'];
};

export type SupportingDoc = {
  __typename?: 'SupportingDoc';
  created_at: Scalars['Time']['output'];
  govt_id: Scalars['String']['output'];
  id: Scalars['UUID']['output'];
  updated_at: Scalars['Time']['output'];
  verified: Verification;
};

export type Title = {
  __typename?: 'Title';
  created_at: Scalars['Time']['output'];
  id: Scalars['UUID']['output'];
  title: Scalars['String']['output'];
  updated_at: Scalars['Time']['output'];
  verified: Verification;
};

export type UpdateTitleVerificationInput = {
  id: Scalars['UUID']['input'];
  verification: Verification;
};

export type User = {
  __typename?: 'User';
  created_at: Scalars['Time']['output'];
  email: Scalars['String']['output'];
  firstname?: Maybe<Scalars['String']['output']>;
  id: Scalars['UUID']['output'];
  lastname?: Maybe<Scalars['String']['output']>;
  updated_at: Scalars['Time']['output'];
  wallet_address: Scalars['String']['output'];
};

export enum Verification {
  Onboarding = 'ONBOARDING',
  Rejected = 'REJECTED',
  Verified = 'VERIFIED'
}

export class TypedDocumentString<TResult, TVariables>
  extends String
  implements DocumentTypeDecoration<TResult, TVariables>
{
  __apiType?: DocumentTypeDecoration<TResult, TVariables>['__apiType'];

  constructor(private value: string, public __meta__?: Record<string, any> | undefined) {
    super(value);
  }

  toString(): string & DocumentTypeDecoration<TResult, TVariables> {
    return this.value;
  }
}
