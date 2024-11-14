package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.twotone.ArrowBack
import androidx.compose.material3.Button
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalSoftwareKeyboardController
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import coil.decode.SvgDecoder
import coil.request.ImageRequest
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation
import com.lomolo.copodapp.ui.viewmodels.ChargingMpesa
import com.lomolo.copodapp.ui.viewmodels.MainViewModel
import com.lomolo.copodapp.ui.viewmodels.MpesaViewModel

object MpesaScreenDestination : Navigation {
    override val title = R.string.mpesa
    override val route = "mpesa"
    const val LAND_TITLE_ID_ARG = "uploadId"
    val routeWithArgs = "$route/{$LAND_TITLE_ID_ARG}"
}

@Composable
fun MpesaScreen(
    modifier: Modifier = Modifier,
    onGoBack: () -> Unit,
    onSuccess: () -> Unit,
    viewModel: MpesaViewModel,
    mainViewModel: MainViewModel,
) {
    LaunchedEffect(key1 = viewModel.chargingMpesa) {
        if (viewModel.chargingMpesa is ChargingMpesa.Paid) {
            onSuccess()
        }
    }
    val mpesa by viewModel.mpesa.collectAsState()
    val keyboardController = LocalSoftwareKeyboardController.current
    val userInfo = mainViewModel.userInfo
    val credentials = mainViewModel.credentials
    val deviceDetails by mainViewModel.deviceDetails.collectAsState()
    val isPhoneValid = viewModel.isValidPhone(mpesa, deviceDetails)
    val context = LocalContext.current

    Scaffold(topBar = {
        TopBar(title = {
            Text(stringResource(R.string.mpesa))
        }, navigationIcon = {
            IconButton(
                onClick = onGoBack,
            ) {
                Icon(
                    Icons.AutoMirrored.TwoTone.ArrowBack,
                    contentDescription = stringResource(R.string.go_back),
                )
            }
        })
    }) { innerPadding ->
        Surface(modifier = modifier) {
            Column(
                Modifier
                    .fillMaxSize()
                    .padding(innerPadding)
                    .padding(8.dp),
                verticalArrangement = Arrangement.spacedBy(8.dp),
            ) {
                OutlinedTextField(
                    isError = mpesa.phone.isNotEmpty() && !isPhoneValid,
                    value = mpesa.phone,
                    onValueChange = { viewModel.setPhone(it) },
                    modifier = Modifier.fillMaxWidth(),
                    singleLine = true,
                    label = {
                        Text(stringResource(R.string.phone_number))
                    },
                    keyboardOptions = KeyboardOptions(
                        keyboardType = KeyboardType.Number,
                        imeAction = ImeAction.Done,
                    ),
                    keyboardActions = KeyboardActions(onDone = {
                        keyboardController?.hide()
                        viewModel.chargeMpesa(userInfo?.email!!, credentials!!.address, deviceDetails)
                    }),
                    leadingIcon = {
                        Row(
                            modifier = Modifier.padding(start = 8.dp),
                            verticalAlignment = Alignment.CenterVertically,
                            horizontalArrangement = Arrangement.spacedBy(8.dp),
                        ) {
                            AsyncImage(
                                model = ImageRequest.Builder(context)
                                    .data(deviceDetails.countryFlag)
                                    .decoderFactory(SvgDecoder.Factory())
                                    .build(),
                                contentScale = ContentScale.Crop,
                                placeholder = painterResource(id = R.drawable.loading_img),
                                modifier = Modifier
                                    .size(32.dp),
                                contentDescription = null
                            )
                            Text(
                                deviceDetails.callingCode,
                            )
                        }
                    },
                )
                Button(
                    onClick = { viewModel.chargeMpesa(userInfo?.email!!, credentials!!.address, deviceDetails) },
                    contentPadding = PaddingValues(16.dp),
                    modifier = Modifier.fillMaxWidth(),
                    shape = MaterialTheme.shapes.extraSmall,
                ) {
                    when (viewModel.chargingMpesa) {
                        ChargingMpesa.Success -> Text(
                            stringResource(R.string.pay),
                            style = MaterialTheme.typography.titleMedium,
                        )

                        ChargingMpesa.Paid -> Text(
                            stringResource(R.string.pay),
                            style = MaterialTheme.typography.titleMedium,
                        )

                        ChargingMpesa.Paying -> CircularProgressIndicator(
                            Modifier.size(20.dp),
                            MaterialTheme.colorScheme.onPrimary,
                        )

                        ChargingMpesa.Loading -> CircularProgressIndicator(
                            Modifier.size(20.dp),
                            MaterialTheme.colorScheme.onPrimary,
                        )
                    }
                }
            }
        }
    }
}