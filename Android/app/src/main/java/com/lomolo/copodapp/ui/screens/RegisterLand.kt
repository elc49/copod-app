package com.lomolo.copodapp.ui.screens

import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.PickVisualMediaRequest
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.twotone.ArrowBack
import androidx.compose.material3.Button
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.navigation.Navigation
import com.lomolo.copodapp.ui.viewmodels.RegisterLandViewModel
import com.lomolo.copodapp.ui.viewmodels.UploadingDoc
import kotlinx.coroutines.launch
import org.koin.androidx.compose.koinViewModel

object RegisterLandScreenDestination : Navigation {
    override val title = R.string.register_land
    override val route = "register-land"
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun RegisterLandScreen(
    modifier: Modifier = Modifier,
    onGoBack: () -> Unit,
    viewModel: RegisterLandViewModel = koinViewModel<RegisterLandViewModel>(),
) {
    val images by viewModel.images.collectAsState()
    val landTitle = when (viewModel.uploadingLandDoc) {
        UploadingDoc.Loading -> {
            R.drawable.loading_img
        }

        UploadingDoc.Success -> {
            R.drawable.upload
        }

        else -> {
            R.drawable.upload
        }
    }
    val idDoc = when (viewModel.uploadingGovtId) {
        UploadingDoc.Loading -> {
            R.drawable.loading_img
        }

        UploadingDoc.Success -> {
            R.drawable.upload
        }

        else -> {
            R.drawable.upload
        }
    }
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val pickLandTitleMedia = rememberLauncherForActivityResult(
        contract = ActivityResultContracts.PickVisualMedia()
    ) {
        if (it != null) {
            val stream = context.contentResolver.openInputStream(it)
            if (stream != null) {
                viewModel.uploadLandTitle(stream)
            }
        }
    }
    val pickGovtIdMedia = rememberLauncherForActivityResult(
        contract = ActivityResultContracts.PickVisualMedia()
    ) {
        if (it != null) {
            val stream = context.contentResolver.openInputStream(it)
            if (stream != null) {
                viewModel.uploadGovtIssuedId(stream)
            }
        }
    }

    Scaffold(topBar = {
        TopAppBar(title = {
            Text(stringResource(R.string.register_land))
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
        Column(
            modifier
                .fillMaxSize()
                .padding(innerPadding)
                .padding(8.dp),
            verticalArrangement = Arrangement.SpaceBetween,
        ) {
            Column {
                Text(
                    stringResource(R.string.land_title_document),
                    fontWeight = FontWeight.Bold,
                    style = MaterialTheme.typography.titleLarge,
                )
                Text(stringResource(R.string.govt_issued_title))
            }
            Box(
                Modifier
                    .fillMaxSize()
                    .clickable {
                        if (viewModel.uploadingLandDoc !is UploadingDoc.Loading) {
                            scope.launch {
                                pickLandTitleMedia.launch(
                                    PickVisualMediaRequest(
                                        ActivityResultContracts.PickVisualMedia.ImageOnly
                                    )
                                )
                            }
                        }
                    },
                contentAlignment = Alignment.Center,
            ) {
                AsyncImage(
                    model = ImageRequest.Builder(context)
                        .data(landTitle)
                        .crossfade(true)
                        .build(),
                    modifier = Modifier
                        .size(100.dp)
                        .clip(MaterialTheme.shapes.small),
                    contentDescription = stringResource(R.string.image),
                )
            }
            Column {
                Text(
                    stringResource(R.string.govt_issued_id),
                    fontWeight = FontWeight.Bold,
                    style = MaterialTheme.typography.titleLarge,
                )
                Text(stringResource(R.string.upload_govt_issued_id))
            }
            Box(
                Modifier
                    .fillMaxSize()
                    .clickable {
                        if (viewModel.uploadingGovtId !is UploadingDoc.Loading) {
                            scope.launch {
                                pickGovtIdMedia.launch(
                                    PickVisualMediaRequest(
                                        ActivityResultContracts.PickVisualMedia.ImageOnly
                                    )
                                )
                            }
                        }
                    },
                contentAlignment = Alignment.Center,
            ) {
                AsyncImage(
                    model = ImageRequest.Builder(context)
                        .data(idDoc)
                        .crossfade(true)
                        .build(),
                    modifier = Modifier
                        .fillMaxSize()
                        .size(100.dp)
                        .clip(MaterialTheme.shapes.small),
                    contentDescription = stringResource(R.string.image),
                )
            }
            Button(
                onClick = {},
                shape = MaterialTheme.shapes.extraSmall,
                contentPadding = PaddingValues(16.dp),
                modifier = Modifier.fillMaxWidth(),
            ) {
                Text(
                    stringResource(R.string.submit),
                    style = MaterialTheme.typography.titleMedium,
                )
            }
        }
    }
}